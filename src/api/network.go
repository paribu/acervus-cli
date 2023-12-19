package api

import (
	"encoding/json"
	"net/http"
)

type NetworksResponse []struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (a *ProjectManagerAPI) GetNetworks() (NetworksResponse, error) {
	resp, err := a.makeAuthenticatedAPIRequest(http.MethodGet, endpoints.network.list, nil)
	if err != nil {
		return nil, err
	}

	var networksResp NetworksResponse
	err = json.Unmarshal(resp, &networksResp)
	if err != nil {
		return nil, err
	}

	return networksResp, nil
}
