package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kwoodhouse93/gotrue-go/types"
)

// GET /admin/users/{user_id}/factors
//
// Get a list of factors for a user.
func (c *Client) AdminListUserFactors(req types.AdminListUserFactorsRequest) (*types.AdminListUserFactorsResponse, error) {
	path := fmt.Sprintf("%s/%s/factors", adminUsersPath, req.UserID)

	r, err := c.newRequest(path, http.MethodGet, nil)
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

	var factors []types.Factor
	err = json.NewDecoder(resp.Body).Decode(&factors)
	if err != nil {
		return nil, err
	}

	return &types.AdminListUserFactorsResponse{
		Factors: factors,
	}, nil
}

// PUT /admin/users/{user_id}/factors/{factor_id}
//
// Update a factor for a user.
func (c *Client) AdminUpdateUserFactor(req types.AdminUpdateUserFactorRequest) (*types.AdminUpdateUserFactorResponse, error) {
	if req.FriendlyName == "" {
		return nil, types.ErrInvalidAdminUpdateFactorRequest
	}

	path := fmt.Sprintf("%s/%s/factors/%s", adminUsersPath, req.UserID, req.FactorID)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(path, http.MethodPut, bytes.NewBuffer(body))
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

	var res types.AdminUpdateUserFactorResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
