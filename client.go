package gotrue

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidProjectReference = errors.New("cannot create gotrue client: invalid project reference")
)

type Client struct {
	client  http.Client
	baseURL string
	apiKey  string
	token   string
}

// Set up a new GoTrue client.
//
// projectReference: The project reference is the unique identifier for your
// Supabase project. It can be found in the Supabase dashboard under project
// settings as Reference ID.
//
// apiKey: The API key is used to authenticate requests to the GoTrue server.
// This should be your anon key.
//
// This function does not validate your project reference. Requests will fail
// if you pass in an invalid project reference.
func New(projectReference string, apiKey string) *Client {
	baseURL := fmt.Sprintf("https://%s.supabase.co/auth/v1", projectReference)
	return &Client{
		client:  http.Client{},
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

// By default, the client will use the supabase project reference and assume
// you are connecting to the GoTrue server as part of a supabase project.
// To connect to a GoTrue server hosted elsewhere, you can specify a custom
// URL using this method.
//
// It returns a copy of the client, so only requests made with the returned
// copy will use the new URL.
//
// This method does not validate the URL you pass in.
func (c Client) WithCustomGoTrueURL(url string) *Client {
	return &Client{
		client:  c.client,
		baseURL: url,
		apiKey:  c.apiKey,
		token:   c.token,
	}
}

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
func (c Client) WithToken(token string) *Client {
	return &Client{
		client:  c.client,
		baseURL: c.baseURL,
		apiKey:  c.apiKey,
		token:   token,
	}
}

// WithClient allows you to pass in your own HTTP client.
//
// It returns a copy of the client, so only requests made with the returned
// copy will use the new HTTP client.
func (c Client) WithClient(client http.Client) *Client {
	return &Client{
		client:  client,
		baseURL: c.baseURL,
		apiKey:  c.apiKey,
		token:   c.token,
	}
}
