package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kwoodhouse93/gotrue-go/types"
)

const ssoPath = "/sso"

// POST /sso
//
// Initiate an SSO session with the given provider.
//
// If successful, the server returns a redirect to the provider's authorization
// URL. The client will follow it and return the final response. Should you
// prefer the client not to follow redirects, you can provide a custom HTTP
// client using WithClient(). See the example below.
//
// Example:
//	c := http.Client{
//		CheckRedirect: func(req *http.Request, via []*http.Request) error {
//			return http.ErrUseLastResponse
//		},
//	}
func (c *Client) SSO(req types.SSORequest) (*http.Response, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(http.MethodPost, ssoPath, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return c.client.Do(r)
}
