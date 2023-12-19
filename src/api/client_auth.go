package api

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/config"
)

type AuthAPI struct {
	apiClient
	BaseURL string
}

func NewAuthAPI() *AuthAPI {
	return &AuthAPI{
		BaseURL: config.AuthServiceURL,
	}
}

func (a *AuthAPI) makeAPIRequest(method, path string, reqData RequestData) ([]byte, error) {
	return a.makeRequest(method, fmt.Sprintf("%s/%s", a.BaseURL, path), reqData)
}

func (a *AuthAPI) makeAuthenticatedAPIRequest(method, path string, reqData RequestData) ([]byte, error) {
	return a.makeAuthenticatedRequest(method, fmt.Sprintf("%s/%s", a.BaseURL, path), reqData)
}
