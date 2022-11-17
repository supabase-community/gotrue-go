package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/supabase-community/gotrue-go/types"
)

// This test relies on running with an empty audit log.
// This appears to be fine with the default test settings, but if it becomes an
// issue locally or in CI, we can add a cleanup step to all tests, or rewrite
// to stop looking at specific numbers of log entries.
func TestAdminAudit(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// User must be authenticated first
	_, err := client.AdminAudit(types.AdminAuditRequest{})
	assert.Error(err)

	client := client.WithToken(adminToken())

	// Get audit logs
	resp, err := client.AdminAudit(types.AdminAuditRequest{})
	require.NoError(err)
	assert.NotNil(resp.Logs)

	// Generate some log entries
	for i := 0; i < 10; i++ {
		pass := randomString(10)
		_, err = client.AdminCreateUser(types.AdminCreateUserRequest{
			Email:    randomEmail(),
			Password: &pass,
		})
		require.NoError(err)
	}

	// Get audit logs with pagination
	resp, err = client.AdminAudit(types.AdminAuditRequest{
		Page:    1,
		PerPage: 4,
	})
	require.NoError(err)
	assert.Len(resp.Logs, 4)
	assert.EqualValues(10, resp.TotalCount)
	assert.EqualValues(2, resp.NextPage)
	assert.EqualValues(3, resp.TotalPages)

	resp, err = client.AdminAudit(types.AdminAuditRequest{
		Page:    2,
		PerPage: 4,
	})
	require.NoError(err)
	assert.Len(resp.Logs, 4)
	assert.EqualValues(10, resp.TotalCount)
	assert.EqualValues(3, resp.NextPage)
	assert.EqualValues(3, resp.TotalPages)

	resp, err = client.AdminAudit(types.AdminAuditRequest{
		Page:    3,
		PerPage: 4,
	})
	require.NoError(err)
	assert.Len(resp.Logs, 2)
	assert.EqualValues(10, resp.TotalCount)
	assert.EqualValues(0, resp.NextPage)
	assert.EqualValues(3, resp.TotalPages)

	// Invalid - empty query
	_, err = client.AdminAudit(types.AdminAuditRequest{
		Query: &types.AuditQuery{},
	})
	assert.Error(err)

	// Invalid - invalid query column
	_, err = client.AdminAudit(types.AdminAuditRequest{
		Query: &types.AuditQuery{
			Column: "invalid",
			Value:  "valid",
		},
	})
	assert.Error(err)

	// Invalid - no query value
	_, err = client.AdminAudit(types.AdminAuditRequest{
		Query: &types.AuditQuery{
			Column: types.AuditQueryColumnAction,
		},
	})
	assert.Error(err)

	// Valid query
	resp, err = client.AdminAudit(types.AdminAuditRequest{
		Query: &types.AuditQuery{
			Column: types.AuditQueryColumnAction,
			Value:  "user_signedup",
		},
	})
	require.NoError(err)
	assert.Len(resp.Logs, 10)

	resp, err = client.AdminAudit(types.AdminAuditRequest{
		Query: &types.AuditQuery{
			Column: types.AuditQueryColumnAction,
			Value:  "not user_signedup",
		},
	})
	require.NoError(err)
	assert.Len(resp.Logs, 0)
}
