package endpoints

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const reauthenticatePath = "/reauthenticate"

// GET /reauthenticate
//
// Sends a nonce to the user's email (preferred) or phone. This endpoint
// requires the user to be logged in / authenticated first. The user needs to
// have either an email or phone number for the nonce to be sent successfully.
func (c *Client) Reauthenticate() error {
	r, err := c.newRequest(reauthenticatePath, http.MethodGet, nil)
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
