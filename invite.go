package gotrue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const invitePath = "/invite"

type InviteRequest struct {
	Email string `json:"email"`
}

type InviteResponse struct {
	ID                 string    `json:"id"`
	Email              string    `json:"email"`
	ConfirmationSentAt time.Time `json:"confirmation_sent_at"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	InvitedAt          time.Time `json:"invited_at"`
}

// POST /invite
//
// Invites a new user with an email.
// This endpoint requires the service_role or supabase_admin JWT set using WithToken.
func (c *Client) Invite(req InviteRequest) (*InviteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(invitePath, http.MethodPost, bytes.NewBuffer(body))
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

	var res InviteResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
