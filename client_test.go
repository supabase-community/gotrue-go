package gotrue_test

import (
	"log"
	"testing"

	"github.com/kwoodhouse93/gotrue-go"
)

const projectReference = "project_ref"
const apiKey = "api_key"
const token = "service_role_token"

func TestAll(t *testing.T) {
	c := gotrue.New(
		projectReference,
		apiKey,
	).WithToken(
		token,
	)

	settings, err := c.GetSettings()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", settings)

	// Disabled as not working, and not needed now
	// pass := "test"
	// err = c.CreateAdminUser(gotrue.CreateAdminUserRequest{
	// 	UserID:       uuid.NewString(),
	// 	Role:         "anon",
	// 	Email:        "test@example.com",
	// 	EmailConfirm: true,
	// 	Password:     &pass,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// signupResp, err := c.Signup(gotrue.SignupRequest{
	// 	Email:    "test@example.com",
	// 	Password: "testme",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%+v", signupResp)

	// inviteResp, err := c.Invite(gotrue.InviteRequest{
	// 	Email: "test@example.com",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%+v", inviteResp)
}
