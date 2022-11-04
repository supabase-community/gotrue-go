package gotrue

import (
	"net/http"

	"github.com/kwoodhouse93/gotrue-go/types"
)

// --- API ---

type Client interface {
	// By default, the client will use the supabase project reference and assume
	// you are connecting to the GoTrue server as part of a supabase project.
	// To connect to a GoTrue server hosted elsewhere, you can specify a custom
	// URL using this method.
	//
	// It returns a copy of the client, so only requests made with the returned
	// copy will use the new URL.
	//
	// This method does not validate the URL you pass in.
	WithCustomGoTrueURL(url string) Client
	// WithToken sets the access token to pass when making admin requests that
	// require token authentication.
	//
	// It returns a copy of the client, so only requests made with the returned
	// copy will use the new token.
	//
	// The token can be your service role if running on a server.
	// REMEMBER TO KEEP YOUR SERVICE ROLE TOKEN SECRET!!!
	//
	// A user token can also be used to make requests on behalf of a user. This is
	// usually preferable to using a service role token.
	WithToken(token string) Client
	// WithClient allows you to pass in your own HTTP client.
	//
	// It returns a copy of the client, so only requests made with the returned
	// copy will use the new HTTP client.
	WithClient(client http.Client) Client

	// POST /admin/generate_link
	//
	// Returns the corresponding email action link based on the type specified.
	// Among other things, the response also contains the query params of the action
	// link as separate JSON fields for convenience (along with the email OTP from
	// which the corresponding token is generated).
	//
	// Requires admin token.
	AdminGenerateLink(req types.AdminGenerateLinkRequest) (*types.AdminGenerateLinkResponse, error)
	// POST /admin/users
	//
	// Creates the user based on the user_id specified.
	//
	// Requires admin token.
	AdminCreateUser(req types.AdminCreateUserRequest) (*types.AdminCreateUserResponse, error)
	// GET /admin/users
	//
	// Get a list of users.
	//
	// Requires admin token.
	AdminListUsers() (*types.AdminListUsersResponse, error)

	// GET /health
	//
	// Check the health of the GoTrue server
	HealthCheck() (*types.HealthCheckResponse, error)

	// POST /invite
	//
	// Invites a new user with an email.
	//
	// Requires service_role or admin token.
	Invite(req types.InviteRequest) (*types.InviteResponse, error)

	// POST /logout
	//
	// Logout a user (Requires authentication).
	//
	// This will revoke all refresh tokens for the user. Remember that the JWT
	// tokens will still be valid for stateless auth until they expires.
	Logout() error

	// GET /settings
	//
	// Returns the publicly available settings for this gotrue instance.
	GetSettings() (*types.SettingsResponse, error)

	// POST /signup
	//
	// Register a new user with an email and password.
	Signup(req types.SignupRequest) (*types.SignupResponse, error)

	// Sign in with email and password
	//
	// This is a convenience method that calls Token with the password grant type
	SignInWithEmailPassword(email, password string) (*types.TokenResponse, error)
	// Sign in with phone and password
	//
	// This is a convenience method that calls Token with the password grant type
	SignInWithPhonePassword(phone, password string) (*types.TokenResponse, error)
	// Sign in with refresh token
	//
	// This is a convenience method that calls Token with the refresh_token grant type
	RefreshToken(refreshToken string) (*types.TokenResponse, error)
	// POST /token
	//
	// This is an OAuth2 endpoint that currently implements the password and
	// refresh_token grant types
	Token(req types.TokenRequest) (*types.TokenResponse, error)

	// GET /user
	//
	// Get the JSON object for the logged in user (requires authentication)
	GetUser() (*types.UserResponse, error)
	// PUT /user
	//
	// Update a user (Requires authentication). Apart from changing email/password,
	// this method can be used to set custom user data. Changing the email will
	// result in a magiclink being sent out.
	UpdateUser(req types.UpdateUserRequest) (*types.UpdateUserResponse, error)
}
