package gotrue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var settingsPath = "/settings"

type ExternalProviders struct {
	Apple     bool `json:"apple"`
	Azure     bool `json:"azure"`
	Bitbucket bool `json:"bitbucket"`
	Discord   bool `json:"discord"`
	Facebook  bool `json:"facebook"`
	Github    bool `json:"github"`
	Gitlab    bool `json:"gitlab"`
	Google    bool `json:"google"`
	Keycloak  bool `json:"keycloak"`
	LinkedIn  bool `json:"linkedin"`
	Notion    bool `json:"notion"`
	Spotify   bool `json:"spotify"`
	Slack     bool `json:"slack"`
	Twitch    bool `json:"twitch"`
	Twitter   bool `json:"twitter"`
	WorkOS    bool `json:"workos"`
}

type SettingsResponse struct {
	Autoconfirm   bool              `json:"autoconfirm"`
	DisableSignup bool              `json:"disable_signup"`
	External      ExternalProviders `json:"external"`
}

// GET /settings
//
// Returns the publicly available settings for this gotrue instance.
func (c *Client) GetSettings() (*SettingsResponse, error) {
	r, err := c.newRequest(settingsPath, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fullBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("response status code %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("response status code %d: %s", resp.StatusCode, fullBody)
	}

	var res SettingsResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
