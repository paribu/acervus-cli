package api

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/config"
)

type authAPI struct {
	apiClient
	BaseURL string
}

func NewAuthAPI() *authAPI {
	return &authAPI{
		BaseURL: config.AuthServiceURL,
	}
}

func (a *authAPI) makeAPIRequest(method, path string, reqData RequestData) ([]byte, error) {
	return a.makeRequest(method, fmt.Sprintf("%s/%s", a.BaseURL, path), reqData)
}

func (a *authAPI) makeAuthenticatedAPIRequest(method, path string, reqData RequestData) ([]byte, error) {
	return a.makeAuthenticatedRequest(method, fmt.Sprintf("%s/%s", a.BaseURL, path), reqData)
}
