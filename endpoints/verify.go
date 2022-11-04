package endpoints

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kwoodhouse93/gotrue-go/types"
)

const verifyPath = "/verify"

// Get /verify
//
// Verify a registration or a password recovery. Type can be signup or recovery
// or magiclink or invite and the token is a token returned from either /signup
// or /recover or /magiclink.
func (c *Client) Verify(req types.VerifyRequest) (*types.VerifyResponse, error) {
	r, err := c.newRequest(verifyPath, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	q := r.URL.Query()
	q.Add("type", string(req.Type))
	q.Add("token", req.Token)
	q.Add("redirect_to", req.RedirectTo)
	r.URL.RawQuery = q.Encode()

	// Set up a client that will not follow the redirect.
	noRedirClient := noRedirClient(c.client)

	resp, err := noRedirClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusSeeOther {
		fullBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("response status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("response status code %d: %s", resp.StatusCode, fullBody)
	}

	redirURL := resp.Header.Get("Location")
	if redirURL == "" {
		return nil, fmt.Errorf("no redirect URL found in response")
	}
	u, err := url.Parse(redirURL)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, err
	}

	return &types.VerifyResponse{
		URL:              redirURL,
		Error:            values.Get("error"),
		ErrorCode:        values.Get("error_code"),
		ErrorDescription: values.Get("error_description"),
	}, nil
}
