package gotrue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kwoodhouse93/gotrue-go/types"
)

const tokenPath = "/token"

var ErrInvalidTokenRequest = errors.New("token request is invalid - grant_type must be password or refresh_token, email and password must be provided for grant_type=password, refresh_token must be provided for grant_type=refresh_token")

type TokenRequest struct {
	GrantType string `json:"-"`

	// Email or Phone, and Password, are required if GrantType is 'password'.
	// They must not be provided if GrantType is 'refresh_token'.
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`

	// RefreshToken is required if GrantType is 'refresh_token'.
	// It must not be provided if GrantType is 'password'.
	RefreshToken string `json:"refresh_token,omitempty"`
}

type TokenResponse struct {
	types.Session
}

// Sign in with email and password
//
// This is a convenience method that calls Token with the password grant type
func (c *Client) SignInWithEmailPassword(email, password string) (*TokenResponse, error) {
	return c.Token(TokenRequest{
		GrantType: "password",
		Email:     email,
		Password:  password,
	})
}

// Sign in with phone and password
//
// This is a convenience method that calls Token with the password grant type
func (c *Client) SignInWithPhonePassword(phone, password string) (*TokenResponse, error) {
	return c.Token(TokenRequest{
		GrantType: "password",
		Phone:     phone,
		Password:  password,
	})
}

// Sign in with refresh token
//
// This is a convenience method that calls Token with the refresh_token grant type
func (c *Client) RefreshToken(refreshToken string) (*TokenResponse, error) {
	return c.Token(TokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	})
}

// POST /token
//
// This is an OAuth2 endpoint that currently implements the password and
// refresh_token grant types
func (c *Client) Token(req TokenRequest) (*TokenResponse, error) {
	switch req.GrantType {
	case "password":
		if (req.Email == "" && req.Phone == "") || req.Password == "" || req.RefreshToken != "" {
			return nil, ErrInvalidTokenRequest
		}
	case "refresh_token":
		if req.RefreshToken == "" || req.Email != "" || req.Phone != "" || req.Password != "" {
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
