package gotrue_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go"
)

func TestSignup(t *testing.T) {
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
	assert.InDelta(time.Now().Unix(), user.ConfirmationSentAt.Unix(), float64(time.Second))
	assert.InDelta(time.Now().Unix(), user.CreatedAt.Unix(), float64(time.Second))
	assert.InDelta(time.Now().Unix(), user.UpdatedAt.Unix(), float64(time.Second))
}
