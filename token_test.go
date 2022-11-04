package gotrue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go"
)

func TestToken(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Test login with email
	email := randomEmail()
	password := "password"

	_, err := autoconfirmClient.Signup(gotrue.SignupRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(err)

	token, err := autoconfirmClient.Token(gotrue.TokenRequest{
		GrantType: "password",
		Email:     email,
		Password:  password,
	})
	require.NoError(err)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal(token.TokenType, "bearer")
	assert.Equal(token.ExpiresIn, 3600)

	// Signin with email convenience method
	token, err = autoconfirmClient.SignInWithEmailPassword(email, password)
	require.NoError(err)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal(token.TokenType, "bearer")
	assert.Equal(token.ExpiresIn, 3600)

	// Test login with phone
	phone := randomPhoneNumber()
	password = "password"

	_, err = autoconfirmClient.Signup(gotrue.SignupRequest{
		Phone:    phone,
		Password: password,
	})
	require.NoError(err)

	token, err = autoconfirmClient.Token(gotrue.TokenRequest{
		GrantType: "password",
		Phone:     phone,
		Password:  password,
	})
	require.NoError(err)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal(token.TokenType, "bearer")
	assert.Equal(token.ExpiresIn, 3600)

	// Signin with phone convenience method
	token, err = autoconfirmClient.SignInWithPhonePassword(phone, password)
	require.NoError(err)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal(token.TokenType, "bearer")
	assert.Equal(token.ExpiresIn, 3600)
}
