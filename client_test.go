package gotrue_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"testing"

	backoff "github.com/cenkalti/backoff/v4"
	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/kwoodhouse93/gotrue-go"
)

const (
	projectReference = "project_ref"
	apiKey           = "api_key"
	jwtSecret        = "secret"
)

var (
	// Global client is used for all tests in this package.
	client *gotrue.Client

	// Used to validate UUIDs.
	uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`)
)

// Utility function to generate some random chars
func randomString(n int) string {
	// Using all lower case because email addresses are lowercased by GoTrue.
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func adminToken() string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud":  "admin",
		"sub":  "admin",
		"role": "supabase_admin",
		"exp":  9999999999,
	})
	token, err := t.SignedString([]byte(jwtSecret))
	if err != nil {
		panic(err)
	}
	return token
}

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
// 	// err = c.AdminCreateUser(gotrue.AdminCreateUserRequest{
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
