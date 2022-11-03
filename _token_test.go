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

	// Create a new user
	user, err := client.Signup(gotrue.SignupRequest{
		Email:    "example@test.com",
		Password: "password",
	})
	require.NoError(err)
	assert.Regexp(uuidRegex, user.ID)
	assert.Equal(user.Email, "example@test.com")

	// TODO: Verify the email

	// Login the user
	token, err := client.Token(gotrue.TokenRequest{
		GrantType: "password",
		Email:     "example@test.com",
		Password:  "password",
	})
	require.NoError(err)
	assert.NotEmpty(token.AccessToken)
	assert.NotEmpty(token.RefreshToken)
	assert.Equal(token.TokenType, "bearer")
	assert.Equal(token.ExpiresIn, 3600)

}
