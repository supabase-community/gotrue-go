package gotrue

import (
	"errors"
	"net/http"

	"github.com/kwoodhouse93/gotrue-go/endpoints"
)

var (
	ErrInvalidProjectReference = errors.New("cannot create gotrue client: invalid project reference")
)

var _ Client = &client{}

type client struct {
	*endpoints.Client
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
func New(projectReference string, apiKey string) Client {
	return &client{
		Client: endpoints.New(projectReference, apiKey),
	}
}

func (c client) WithCustomGoTrueURL(url string) Client {
	return &client{
		Client: c.Client.WithCustomGoTrueURL(url),
	}
}

func (c client) WithToken(token string) Client {
	return &client{
		Client: c.Client.WithToken(token),
	}
}

func (c client) WithClient(httpClient http.Client) Client {
	return &client{
		Client: c.Client.WithClient(httpClient),
	}
}
