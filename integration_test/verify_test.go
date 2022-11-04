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

	// Unauthorized client
	resp, err := autoconfirmClient.Verify(types.VerifyRequest{
		Type:       types.VerificationTypeSignup,
		Token:      "abcde",
		RedirectTo: "http://localhost:3000",
	})
	require.NoError(err)
	assert.NotEmpty(resp.URL)
	assert.Equal("401", resp.ErrorCode)
	assert.Equal("unauthorized_client", resp.Error)

	// Authorized client
	// TODO: Test
}
