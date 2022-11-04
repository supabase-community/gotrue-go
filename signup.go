package gotrue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"github.com/kwoodhouse93/gotrue-go/types"
)

const signupPath = "/signup"

type SignupRequest struct {
	Email    string                 `json:"email"`
	Phone    string                 `json:"phone"`
	Password string                 `json:"password"`
	Data     map[string]interface{} `json:"data"`
}

type SignupResponse struct {
	// Response if autoconfirm is off
	types.User

	// Response if autoconfirm is on
	types.Session
}

// POST /signup
//
// Register a new user with an email and password.
func (c *Client) Signup(req SignupRequest) (*SignupResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(signupPath, http.MethodPost, bytes.NewBuffer(body))
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

	var res SignupResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	// To make the response easier to consume, if autoconfirm was enabled, the
	// session user should be popuated. Copy that into the embedded user type
	// so it's easier to access.
	//
	// i.e. we can access user fields like res.Email regardless of whether
	// we got back a session or a user.
	if res.Session.User.ID != uuid.Nil {
		res.User = res.Session.User
	}

	return &res, nil
}
