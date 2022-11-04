package endpoints

import (
	"io"
	"net/http"
)

func (c *Client) newRequest(path string, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("apiKey", c.apiKey)
	if c.token != "" {
		req.Header.Add("Authorization", "Bearer "+c.token)
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, nil
}
