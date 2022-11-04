package integration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettings(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	settings, err := client.GetSettings()
	require.NoError(err)
	assert.Equal(settings.Autoconfirm, false)
	assert.Equal(settings.DisableSignup, false)
}
