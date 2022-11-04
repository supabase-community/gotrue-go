package gotrue_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go"
)

func TestSignup(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Signup with email
	email := randomEmail()
	userResp, err := client.Signup(gotrue.SignupRequest{
		Email:    email,
		Password: "password",
	})
	require.NoError(err)
	assert.NotEqual(uuid.UUID{}, userResp.ID)
	assert.Equal(userResp.Email, email)
	assert.InDelta(time.Now().Unix(), userResp.ConfirmationSentAt.Unix(), float64(time.Second))
	assert.InDelta(time.Now().Unix(), userResp.CreatedAt.Unix(), float64(time.Second))
	assert.InDelta(time.Now().Unix(), userResp.UpdatedAt.Unix(), float64(time.Second))

	// Duplicate signup
	dupeUserResp, err := client.Signup(gotrue.SignupRequest{
		Email:    email,
		Password: "password",
	})
	require.NoError(err)
	assert.NotEqual(uuid.UUID{}, dupeUserResp.ID)
	assert.Equal(dupeUserResp.Email, email)
	assert.InDelta(time.Now().Unix(), dupeUserResp.ConfirmationSentAt.Unix(), float64(time.Second))
	assert.InDelta(time.Now().Unix(), dupeUserResp.CreatedAt.Unix(), float64(time.Second))
	assert.InDelta(time.Now().Unix(), dupeUserResp.UpdatedAt.Unix(), float64(time.Second))
	assert.Equal(userResp.ID, dupeUserResp.ID)

	// Sign up with phone
	// Will error because SMS is not configured on the test server.
	user, err := client.Signup(gotrue.SignupRequest{
		Phone:    "+15555555555",
		Password: "password",
	})
	assert.Error(err)
	assert.Nil(user)
}
