package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/supabase-community/gotrue-go/types"
)

func TestAuthorize(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Provider not enabled
	_, err := client.Authorize(types.AuthorizeRequest{
		Provider: "apple",
	})
	assert.Error(err)

	// Provider enabled
	resp, err := autoconfirmClient.Authorize(types.AuthorizeRequest{
		Provider: "github",
	})
	require.NoError(err)
	assert.Contains(resp.AuthorizationURL, "github.com/login/oauth/authorize")

	// Test login with PKCE
	request := types.AuthorizeRequest{
		Provider: "google",
		FlowType: "pkce",
	}
	resp, err = client.Authorize(request)
	require.NoError(err)
	require.NotEmpty(resp.AuthorizationURL)
	require.NotEmpty(resp.Verifier)

	// No provider chosen
	_, err = autoconfirmClient.Authorize(types.AuthorizeRequest{})
	assert.Error(err)
}
