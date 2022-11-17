package endpoints

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/supabase-community/gotrue-go/types"
)

const authorizePath = "/authorize"

// GET /authorize
//
// Get access_token from external oauth provider.
//
// Scopes are optional additional scopes depending on the provider (email and
// name are requested by default).
//
// If successful, the server returns a redirect response. This method will not
// follow the redirect, but instead returns the URL the client was told to
// redirect to.
func (c *Client) Authorize(req types.AuthorizeRequest) (*types.AuthorizeResponse, error) {
	r, err := c.newRequest(authorizePath, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	q := r.URL.Query()
	q.Add("provider", string(req.Provider))
	q.Add("scopes", req.Scopes)
	r.URL.RawQuery = q.Encode()

	// Set up a client that will not follow the redirect.
	noRedirClient := noRedirClient(c.client)

	resp, err := noRedirClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		fullBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("response status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("response status code %d: %s", resp.StatusCode, fullBody)
	}

	url := resp.Header.Get("Location")
	if url == "" {
		return nil, fmt.Errorf("no redirect URL found in response")
	}
	return &types.AuthorizeResponse{
		AuthorizationURL: url,
	}, nil
}
