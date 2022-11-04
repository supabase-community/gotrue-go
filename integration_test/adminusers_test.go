package integration_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go/types"
)

func TestAdminCreateUser(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	admin := withAdmin(client)

	pass := "password"
	email := randomEmail()
	req := types.AdminCreateUserRequest{
		Email:    email,
		Role:     "admin",
		Password: &pass,
	}
	resp, err := admin.AdminCreateUser(req)
	require.NoError(err)
	require.Regexp(uuidRegex, resp.ID)
	assert.Equal(resp.Email, email)
	assert.Equal(resp.Role, "admin")
}

func TestAdminListUsers(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	admin := withAdmin(client)

	// Create a user that we know should be returned
	pass := "password"
	email := randomEmail()
	req := types.AdminCreateUserRequest{
		Email:    email,
		Role:     "test",
		Password: &pass,
	}
	createResp, err := admin.AdminCreateUser(req)
	require.NoError(err)
	require.Regexp(uuidRegex, createResp.ID)

	// Then list and look up the user we just created
	resp, err := admin.AdminListUsers()
	require.NoError(err)
	assert.NotEmpty(resp)
	for _, u := range resp.Users {
		assert.NotEqual(uuid.Nil, u.ID)
		if u.ID == createResp.ID {
			assert.Equal(u.Email, createResp.Email)
		}
	}
}
