package gotrue_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go"
)

func TestAdminGenerateLink(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	c := client.WithToken(adminToken())

	email := randomEmail()
	resp, err := c.AdminGenerateLink(gotrue.AdminGenerateLinkRequest{
		Type:       "signup",
		Email:      email,
		Password:   "password",
		Data:       map[string]interface{}{},
		RedirectTo: "http://localhost:3000",
	})
	require.NoError(err)
	assert.EqualValues(resp.VerificationType, "signup")
	linkRegexp := regexp.MustCompile(`^http://localhost:9999/\?token=[a-zA-Z0-9_-]+&type=signup&redirect_to=http://localhost:3000$`)
	assert.Regexp(linkRegexp, resp.ActionLink)

	// TODO: Test the rest of the response
	// TODO: Test invalid requests
	// TODO: Test updating emails with generate link
}
