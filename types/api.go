package types

import (
	"errors"
	"fmt"
)

// --- Errors ---

type ErrInvalidGenerateLinkRequest struct {
	message string
}

func NewErrInvalidGenerateLinkRequest(message string) *ErrInvalidGenerateLinkRequest {
	return &ErrInvalidGenerateLinkRequest{message: message}
}

func (e *ErrInvalidGenerateLinkRequest) Error() string {
	return fmt.Sprintf("generate link request is invalid - %s", e.message)
}

var (
	ErrInvalidTokenRequest  = errors.New("token request is invalid - grant_type must be password or refresh_token, email and password must be provided for grant_type=password, refresh_token must be provided for grant_type=refresh_token")
	ErrInvalidVerifyRequest = errors.New("verify request is invalid - type, token and redirect_to must be provided, and email or phone must be provided to VerifyForUser")
)

// --- Request/Response Types ---

type LinkType string

const (
	LinkTypeSignup             LinkType = "signup"
	LinkTypeMagicLink          LinkType = "magiclink"
	LinkTypeRecovery           LinkType = "recovery"
	LinkTypeInvite             LinkType = "invite"
	LinkTypeEmailChangeCurrent LinkType = "email_change_current"
	LinkTypeEmailChangeNew     LinkType = "email_change_new"
)

type AdminGenerateLinkRequest struct {
	Type       LinkType               `json:"type"`
	Email      string                 `json:"email"`
	NewEmail   string                 `json:"new_email"`
	Password   string                 `json:"password"`
	Data       map[string]interface{} `json:"data"`
	RedirectTo string                 `json:"redirect_to"`
}

type AdminGenerateLinkResponse struct {
	ActionLink       string   `json:"action_link"`
	EmailOTP         string   `json:"email_otp"`
	HashedToken      string   `json:"hashed_token"`
	RedirectTo       string   `json:"redirect_to"`
	VerificationType LinkType `json:"verification_type"`

	User
}

type AdminCreateUserRequest struct {
	Aud          string                 `json:"aud"`
	Role         string                 `json:"role"`
	Email        string                 `json:"email"`
	Phone        string                 `json:"phone"`
	Password     *string                `json:"password"` // Only if type = signup
	EmailConfirm bool                   `json:"email_confirm"`
	PhoneConfirm bool                   `json:"phone_confirm"`
	UserMetadata map[string]interface{} `json:"user_metadata"`
	AppMetadata  map[string]interface{} `json:"app_metadata"`
}

type AdminCreateUserResponse struct {
	User
}

type AdminListUsersResponse struct {
	Users []User `json:"users"`
}

type Provider string

const (
	ProviderApple     Provider = "apple"
	ProviderAzure     Provider = "azure"
	ProviderBitbucket Provider = "bitbucket"
	ProviderDiscord   Provider = "discord"
	ProviderGitHub    Provider = "github"
	ProviderGitLab    Provider = "gitlab"
	ProviderGoogle    Provider = "google"
	ProviderKeycloak  Provider = "keycloak"
	ProviderLinkedin  Provider = "linkedin"
	ProviderFacebook  Provider = "facebook"
	ProviderNotion    Provider = "notion"
	ProviderSpotify   Provider = "spotify"
	ProviderSlack     Provider = "slack"
	ProviderTwitch    Provider = "twitch"
	ProviderTwitter   Provider = "twitter"
	ProviderWorkOS    Provider = "workos"
	ProviderZoom      Provider = "zoom"
)

type AuthorizeRequest struct {
	Provider Provider
	Scopes   string
}

type AuthorizeResponse struct {
	AuthorizationURL string
}

type HealthCheckResponse struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type InviteRequest struct {
	Email string                 `json:"email"`
	Data  map[string]interface{} `json:"data"`
}

type InviteResponse struct {
	User
}

// DEPRECATED: Use /otp with Email and CreateUser=true instead of /magiclink.
type MagiclinkRequest struct {
	Email string `json:"email"`
}

type OTPRequest struct {
	Email      string                 `json:"email"`
	Phone      string                 `json:"phone"`
	CreateUser bool                   `json:"create_user"`
	Data       map[string]interface{} `json:"data"`
}

type RecoverRequest struct {
	Email string `json:"email"`
}

type ExternalProviders struct {
	Apple     bool `json:"apple"`
	Azure     bool `json:"azure"`
	Bitbucket bool `json:"bitbucket"`
	Discord   bool `json:"discord"`
	Email     bool `json:"email"`
	Facebook  bool `json:"facebook"`
	GitHub    bool `json:"github"`
	GitLab    bool `json:"gitlab"`
	Google    bool `json:"google"`
	Keycloak  bool `json:"keycloak"`
	Linkedin  bool `json:"linkedin"`
	Notion    bool `json:"notion"`
	Phone     bool `json:"phone"`
	SAML      bool `json:"saml"`
	Slack     bool `json:"slack"`
	Spotify   bool `json:"spotify"`
	Twitch    bool `json:"twitch"`
	Twitter   bool `json:"twitter"`
	WorkOS    bool `json:"workos"`
	Zoom      bool `json:"zoom"`
}

type SettingsResponse struct {
	DisableSignup     bool              `json:"disable_signup"`
	Autoconfirm       bool              `json:"autoconfirm"`
	MailerAutoconfirm bool              `json:"mailer_autoconfirm"`
	PhoneAutoconfirm  bool              `json:"phone_autoconfirm"`
	SmsProvider       string            `json:"sms_provider"`
	MFAEnabled        bool              `json:"mfa_enabled"`
	External          ExternalProviders `json:"external"`
}

type SignupRequest struct {
	Email    string                 `json:"email"`
	Phone    string                 `json:"phone"`
	Password string                 `json:"password"`
	Data     map[string]interface{} `json:"data"`
}

type SignupResponse struct {
	// Response if autoconfirm is off
	User

	// Response if autoconfirm is on
	Session
}

type TokenRequest struct {
	GrantType string `json:"-"`

	// Email or Phone, and Password, are required if GrantType is 'password'.
	// They must not be provided if GrantType is 'refresh_token'.
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`

	// RefreshToken is required if GrantType is 'refresh_token'.
	// It must not be provided if GrantType is 'password'.
	RefreshToken string `json:"refresh_token,omitempty"`
}

type TokenResponse struct {
	Session
}

type UserResponse struct {
	User
}

type UpdateUserRequest struct {
	Email    string                 `json:"email"`
	Password *string                `json:"password"`
	Nonce    string                 `json:"nonce"`
	Data     map[string]interface{} `json:"data"`
	AppData  map[string]interface{} `json:"app_metadata,omitempty"`
	Phone    string                 `json:"phone"`
}

type UpdateUserResponse struct {
	User
}

type VerificationType string

const (
	VerificationTypeSignup      = "signup"
	VerificationTypeRecovery    = "recovery"
	VerificationTypeInvite      = "invite"
	VerificationTypeMagiclink   = "magiclink"
	VerificationTypeEmailChange = "email_change"
	VerificationTypeSMS         = "sms"
	VerificationTypePhoneChange = "phone_change"
)

type VerifyRequest struct {
	Type       VerificationType
	Token      string
	RedirectTo string
}

type VerifyResponse struct {
	URL string

	// The fields below are returned only for a successful response.
	AccessToken  string
	TokenType    string
	ExpiresIn    int
	RefreshToken string
	Type         VerificationType

	// The fields below are returned if there was an error verifying.
	Error            string
	ErrorCode        string
	ErrorDescription string
}

type VerifyForUserRequest struct {
	Type       VerificationType `json:"type"`
	Token      string           `json:"token"`
	RedirectTo string           `json:"redirect_to"`
	Email      string           `json:"email"`
	Phone      string           `json:"phone"`
}

type VerifyForUserResponse struct {
	Session
}
