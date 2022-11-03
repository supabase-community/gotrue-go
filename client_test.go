package gotrue_test

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	backoff "github.com/cenkalti/backoff/v4"

	"github.com/kwoodhouse93/gotrue-go"
)

const (
	projectReference = "project_ref"
	apiKey           = "api_key"
)

var (
	// Global client is used for all tests in this package.
	client *gotrue.Client

	// Used to validate UUIDs.
	uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`)
)

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
// 	// 	Password: "test me",
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
