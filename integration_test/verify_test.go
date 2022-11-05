package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go/types"
)

func TestVerify(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	client := autoconfirmClient

	// Test Verify, unauthorized
	vResp, err := client.Verify(types.VerifyRequest{
		Type:       types.VerificationTypeSignup,
		Token:      "abcde",
		RedirectTo: "http://localhost:3000",
	})
	require.NoError(err)
	assert.NotEmpty(vResp.URL)
	assert.Equal("401", vResp.ErrorCode)
	assert.Equal("unauthorized_client", vResp.Error)

	// Test Verify, invalid request
	_, err = client.Verify(types.VerifyRequest{})
	assert.Error(err)

	// Test VerifyForUser, user doesn't exist
	_, err = client.VerifyForUser(types.VerifyForUserRequest{
		Type:       types.VerificationTypeInvite,
		Token:      "abcde",
		RedirectTo: "http://localhost:3000",
		Email:      randomEmail(),
	})
	assert.Error(err)

	// Test VerifyForUser, user exists, unauthorized
	email := randomEmail()
	_, err = client.Signup(types.SignupRequest{
		Email:    email,
		Password: "password",
	})
	require.NoError(err)

	_, err = client.VerifyForUser(types.VerifyForUserRequest{
		Type:       types.VerificationTypeSignup,
		Token:      "abcde",
		RedirectTo: "http://localhost:3000",
		Email:      email,
	})
	assert.Error(err)

	// Test VerifyForUser, invalid request
	_, err = client.VerifyForUser(types.VerifyForUserRequest{})
	assert.Error(err)
}
