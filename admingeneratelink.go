package gotrue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const adminGenerateLinkPath = "/admin/generate_link"

type ErrInvalidGenerateLinkRequest struct {
	message string
}

func (e *ErrInvalidGenerateLinkRequest) Error() string {
	return fmt.Sprintf("generate link request is invalid - %s", e.message)
}

type LinkType string

const (
	LinkTypeSignup             LinkType = "signup"
	LinkTypeMagicLink          LinkType = "magiclink"
	LinkTypeRecovery           LinkType = "recovery"
	LinkTypeInvite             LinkType = "invite"
	LinkTypeEmailChangeCurrent LinkType = "email_change_current"
	LinkTypeEmailChangeNew     LinkType = "email_change_new"
)

type AdminGenerateLinkRequest struct {
	Type       LinkType               `json:"type"`
	Email      string                 `json:"email"`
	NewEmail   string                 `json:"new_email"`
	Password   string                 `json:"password"`
	Data       map[string]interface{} `json:"data"`
	RedirectTo string                 `json:"redirect_to"`
}

type AdminGenerateLinkResponse struct {
	ActionLink       string   `json:"action_link"`
	EmailOTP         string   `json:"email_otp"`
	HashedToken      string   `json:"hashed_token"`
	RedirectTo       string   `json:"redirect_to"`
	VerificationType LinkType `json:"verification_type"`
}

func validateAdminGenerateLinkRequest(req AdminGenerateLinkRequest) error {
	switch req.Type {
	case LinkTypeSignup:
		if req.Email == "" || req.Password == "" {
			return &ErrInvalidGenerateLinkRequest{
				message: "email and password must be provided if Type is signup",
			}
		}
	case LinkTypeMagicLink, LinkTypeInvite:
		if req.Email == "" {
			return &ErrInvalidGenerateLinkRequest{
				message: "email must be provided if Type is magiclink or invite",
			}
		}
	case LinkTypeRecovery:
		if req.Email == "" {
			return &ErrInvalidGenerateLinkRequest{
				message: "email must be provided if Type is recovery",
			}
		}
		if len(req.Data) > 0 {
			return &ErrInvalidGenerateLinkRequest{
				message: "data must not be provided if Type is recovery",
			}
		}
	case LinkTypeEmailChangeCurrent, LinkTypeEmailChangeNew:
		if req.Email == "" || req.NewEmail == "" {
			return &ErrInvalidGenerateLinkRequest{
				message: "email and new_email must be provided if Type is email_change_current or email_change_new",
			}
		}
		if len(req.Data) > 0 {
			return &ErrInvalidGenerateLinkRequest{
				message: "data must not be provided if Type is email_change_current or email_change_new",
			}
		}
	}

	return nil
}

// POST /admin/generate_link
//
// Returns the corresponding email action link based on the type specified.
// Among other things, the response also contains the query params of the action
// link as separate JSON fields for convenience (along with the email OTP from
// which the corresponding token is generated).
func (c *Client) AdminGenerateLink(req AdminGenerateLinkRequest) (*AdminGenerateLinkResponse, error) {
	err := validateAdminGenerateLinkRequest(req)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(adminGenerateLinkPath, http.MethodPost, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fullBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("response status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("response status code %d: %s", resp.StatusCode, fullBody)
	}

	var res AdminGenerateLinkResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
