package gotrue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go"
)

func TestAdminGenerateLink(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	c := client.WithToken(adminToken)

	resp, err := c.AdminGenerateLink(gotrue.AdminGenerateLinkRequest{
		Type:       "signup",
		Email:      "example@test.com",
		Password:   "password",
		Data:       map[string]interface{}{},
		RedirectTo: "http://localhost:3000",
	})
	require.NoError(err)
	assert.Equal(resp.VerificationType, "signup")
}
