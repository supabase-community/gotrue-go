package integration_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go/types"
)

func TestAdminListUserFactors(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Need admin credential
	_, err := client.AdminListUserFactors(types.AdminListUserFactorsRequest{
		UserID: uuid.Nil,
	})
	assert.Error(err)

	admin := withAdmin(client)

	// Cannot get a user that doesn't exist
	_, err = admin.AdminListUserFactors(types.AdminListUserFactorsRequest{
		UserID: uuid.New(),
	})
	assert.Error(err)

	// Create a user
	email := randomEmail()
	session, err := autoconfirmClient.Signup(types.SignupRequest{
		Email:    email,
		Password: "password",
	})
	require.NoError(err)
	require.Regexp(uuidRegex, session.User.ID)

	// Get that user
	resp, err := admin.AdminListUserFactors(types.AdminListUserFactorsRequest{
		UserID: session.User.ID,
	})
	require.NoError(err)
	assert.Len(resp.Factors, 0)

	// Enroll factor
	enrollResp, err := autoconfirmClient.WithToken(session.AccessToken).EnrollFactor(types.EnrollFactorRequest{
		FactorType:   types.FactorTypeTOTP,
		FriendlyName: "Test Factor",
		Issuer:       "example.com",
	})
	require.NoError(err)
	assert.Equal(types.FactorTypeTOTP, enrollResp.Type)
	assert.NotEqual(uuid.Nil, enrollResp.ID)

	// Get that user again
	// Get that user
	resp, err = admin.AdminListUserFactors(types.AdminListUserFactorsRequest{
		UserID: session.User.ID,
	})
	require.NoError(err)
	require.Len(resp.Factors, 1)
	assert.Equal(resp.Factors[0].ID, enrollResp.ID)
}
