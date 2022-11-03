package gotrue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const tokenPath = "/token"

var ErrInvalidTokenRequest = errors.New("token request is invalid - grant_type must be password or refresh_token, email and password must be provided for grant_type=password, refresh_token must be provided for grant_type=refresh_token")

type TokenRequest struct {
	GrantType string `json:"-"`

	// Email and Password are required if GrantType is 'password'.
	// They must not be provided if GrantType is 'refresh_token'.
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`

	// RefreshToken is required if GrantType is 'refresh_token'.
	// It must not be provided if GrantType is 'password'.
	RefreshToken string `json:"refresh_token,omitempty"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// POST /token
//
// This is an OAuth2 endpoint that currently implements the password and
// refresh_token grant types
func (c *Client) Token(req TokenRequest) (*TokenResponse, error) {
	switch req.GrantType {
	case "password":
		if req.Email == "" || req.Password == "" || req.RefreshToken != "" {
			return nil, ErrInvalidTokenRequest
		}
	case "refresh_token":
		if req.RefreshToken == "" || req.Email != "" || req.Password != "" {
			return nil, ErrInvalidTokenRequest
		}
	default:
		return nil, ErrInvalidTokenRequest
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	r, err := c.newRequest(tokenPath+"?grant_type="+req.GrantType, http.MethodPost, bytes.NewBuffer(body))
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

	var res TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
