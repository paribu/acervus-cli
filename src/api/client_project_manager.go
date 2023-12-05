package api

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/config"
)

type projectManagerAPI struct {
	apiClient
	BaseURL string
}

func NewProjectManagerAPI() *projectManagerAPI {
	return &projectManagerAPI{
		BaseURL: config.ProjectManagerServiceURL,
	}
}

func (a *projectManagerAPI) makeAuthenticatedAPIRequest(method, path string, body []byte) ([]byte, error) {
	return a.makeAuthenticatedRequest(method, fmt.Sprintf("%s/%s", a.BaseURL, path), body)
}
