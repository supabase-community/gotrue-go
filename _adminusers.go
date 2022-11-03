package gotrue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kwoodhouse93/nav-battle/pkg/supabase-gotrue-go/types"
)

const adminUsersPath = "/admin/users/"

type CreateAdminUserRequest struct {
	UserID string `json:"-"`

	Role         string                 `json:"role"`
	Email        string                 `json:"email"`
	Phone        string                 `json:"phone"`
	Password     *string                `json:"password,omitempty"` // Only if type = signup
	EmailConfirm bool                   `json:"email_confirm"`
	PhoneConfirm bool                   `json:"phone_confirm"`
	UserMetadata map[string]interface{} `json:"user_metadata"`
	AppMetadata  map[string]interface{} `json:"app_metadata"`
	BanDuration  types.BanDuration      `json:"ban_duration"`
}

// POST /admin/users/<user_id>
//
// Creates the user based on the user_id specified.
func (c *Client) CreateAdminUser(req CreateAdminUserRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r, err := c.newRequest(adminUsersPath+req.UserID, http.MethodPost, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fullBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("response status code %d", resp.StatusCode)
		}
		return fmt.Errorf("response status code %d: %s", resp.StatusCode, fullBody)
	}

	return nil
}
