package gotrue_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kwoodhouse93/gotrue-go"
)

const projectReference = "project_ref"
const apiKey = "api_key"
const token = "service_role_token"

var client *gotrue.Client

func TestMain(m *testing.M) {
	client = gotrue.New(projectReference, apiKey).WithCustomGoTrueURL("http://localhost:9999")

	// Ensure the server is ready before running tests.
	err := backoff.Retry(
		func() error {
			health, err := client.HealthCheck()
			if err != nil {
				return err
			}
			if health.Name != "GoTrue" {
				return fmt.Errorf("health check - unexpected server name: %s", health.Name)
			}
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 10),
	)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestHealth(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	client := gotrue.New(projectReference, apiKey).WithCustomGoTrueURL("http://localhost:9999")
	health, err := client.HealthCheck()
	require.NoError(err)
	assert.Equal(health.Name, "GoTrue")
}

// func TestAll(t *testing.T) {
// 	c := client.WithToken(
// 		token,
// 	)

// 	settings, err := c.GetSettings()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%+v", settings)

// 	// Disabled as not working, and not needed now
// 	// pass := "test"
// 	// err = c.CreateAdminUser(gotrue.CreateAdminUserRequest{
// 	// 	UserID:       uuid.NewString(),
// 	// 	Role:         "anon",
// 	// 	Email:        "test@example.com",
// 	// 	EmailConfirm: true,
// 	// 	Password:     &pass,
// 	// })
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// signupResp, err := c.Signup(gotrue.SignupRequest{
// 	// 	Email:    "test@example.com",
// 	// 	Password: "testme",
// 	// })
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// log.Printf("%+v", signupResp)

// 	// inviteResp, err := c.Invite(gotrue.InviteRequest{
// 	// 	Email: "test@example.com",
// 	// })
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// log.Printf("%+v", inviteResp)
// }
