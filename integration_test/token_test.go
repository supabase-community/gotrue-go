package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go/types"
)

func TestToken(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Test login with email
	email := randomEmail()
	password := "password"

	_, err := autoconfirmClient.Signup(types.SignupRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(err)

	token, err := autoconfirmClient.Token(types.TokenRequest{
		GrantType: "password",
		Email:     email,
		Password:  password,
	})
	require.NoError(err)
	assert.Equal(email, token.User.Email)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal("bearer", token.TokenType)
	assert.Equal(3600, token.ExpiresIn)

	// Signin with email convenience method
	token, err = autoconfirmClient.SignInWithEmailPassword(email, password)
	require.NoError(err)
	assert.Equal(email, token.User.Email)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal("bearer", token.TokenType)
	assert.Equal(3600, token.ExpiresIn)

	// Test login with phone
	phone := randomPhoneNumber()
	password = "password"

	_, err = autoconfirmClient.Signup(types.SignupRequest{
		Phone:    phone,
		Password: password,
	})
	require.NoError(err)

	token, err = autoconfirmClient.Token(types.TokenRequest{
		GrantType: "password",
		Phone:     phone,
		Password:  password,
	})
	require.NoError(err)
	assert.Equal(phone, token.User.Phone)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal("bearer", token.TokenType)
	assert.Equal(3600, token.ExpiresIn)

	// Signin with phone convenience method
	token, err = autoconfirmClient.SignInWithPhonePassword(phone, password)
	require.NoError(err)
	assert.Equal(phone, token.User.Phone)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal("bearer", token.TokenType)
	assert.Equal(3600, token.ExpiresIn)

	// Test login with refresh token
	email = randomEmail()
	user, err := autoconfirmClient.Signup(types.SignupRequest{
		Email:    email,
		Password: "password",
	})
	require.NoError(err)
	require.NotEmpty(user.RefreshToken)

	token, err = autoconfirmClient.Token(types.TokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: user.RefreshToken,
	})
	require.NoError(err)
	assert.Equal(email, token.User.Email)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal("bearer", token.TokenType)
	assert.Equal(3600, token.ExpiresIn)

	// Refresh token convenience method
	token, err = autoconfirmClient.RefreshToken(token.RefreshToken)
	require.NoError(err)
	assert.Equal(email, token.User.Email)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal("bearer", token.TokenType)
	assert.Equal(3600, token.ExpiresIn)

	// Invalid input tests
	tests := map[string]types.TokenRequest{
		"invalid_grant_type": {
			GrantType: "invalid",
		},
		"password/missing_email_phone": {
			GrantType: "password",
		},
		"password/missing_password": {
			GrantType: "password",
			Email:     email,
		},
		"password/refresh_token_provided": {
			GrantType:    "password",
			Email:        email,
			Password:     password,
			RefreshToken: "refresh_token",
		},
		"refresh_token/missing_refresh_token": {
			GrantType: "refresh_token",
		},
		"refresh_token/email_provided": {
			GrantType: "refresh_token",
			Email:     email,
		},
		"refresh_token/phone_provided": {
			GrantType: "refresh_token",
			Phone:     phone,
		},
		"refresh_token/password_provided": {
			GrantType: "refresh_token",
			Password:  password,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := autoconfirmClient.Token(test)
			require.Error(err)
		})
	}
}
