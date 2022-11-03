package gotrue_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go"
)

func TestAdminCreateUser(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	c := client.WithToken(adminToken())

	pass := "password"
	email := fmt.Sprintf("%s@test.com", randomString(10))
	req := gotrue.AdminCreateUserRequest{
		Email:    email,
		Role:     "admin",
		Password: &pass,
	}
	resp, err := c.AdminCreateUser(req)
	require.NoError(err)
	require.Regexp(uuidRegex, resp.ID)
	assert.Equal(resp.Email, email)
	assert.Equal(resp.Role, "admin")
}

func TestAdminListUsers(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	c := client.WithToken(adminToken())

	// Create a user that we know should be returned
	pass := "password"
	email := fmt.Sprintf("%s@test.com", randomString(10))
	req := gotrue.AdminCreateUserRequest{
		Email:    email,
		Role:     "test",
		Password: &pass,
	}
	createResp, err := c.AdminCreateUser(req)
	require.NoError(err)
	require.Regexp(uuidRegex, createResp.ID)

	// Then list and look up the user we just created
	resp, err := c.AdminListUsers()
	require.NoError(err)
	assert.NotEmpty(resp)
	for _, u := range resp.Users {
		assert.Regexp(uuidRegex, u.ID)
		if u.ID == createResp.ID {
			assert.Equal(u.Email, createResp.Email)
		}
	}
}
