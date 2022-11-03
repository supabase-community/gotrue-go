package gotrue_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go"
)

func TestCreateAdminUser(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	c := client.WithToken(adminToken())

	pass := "password"
	email := fmt.Sprintf("%s@test.com", randomString(10))
	req := gotrue.CreateAdminUserRequest{
		Email:    email,
		Role:     "admin",
		Password: &pass,
	}
	resp, err := c.CreateAdminUser(req)
	require.NoError(err)
	assert.Regexp(uuidRegex, resp.ID)
	assert.Equal(resp.Email, email)
	assert.Equal(resp.Role, "admin")
}
