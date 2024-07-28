package endpoints

import (
	"fmt"
	"io"
	"net/http"
)

const logoutPath = "/logout"

// POST /logout
//
// Logout a user (Requires authentication).
//
// This will revoke all refresh tokens for the user. Remember that the JWT
// tokens will still be valid for stateless auth until they expires.
func (c *Client) Logout() error {
	r, err := c.newRequest(logoutPath, http.MethodPost, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fullBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("response status code %d", resp.StatusCode)
		}
		return fmt.Errorf("response status code %d: %s", resp.StatusCode, fullBody)
	}

	return nil
}
