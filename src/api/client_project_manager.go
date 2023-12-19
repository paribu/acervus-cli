package api

import (
	"fmt"

	"github.com/paribu/acervus-cli/src/config"
)

type ProjectManagerAPI struct {
	apiClient
	BaseURL string
}

func NewProjectManagerAPI() *ProjectManagerAPI {
	return &ProjectManagerAPI{
		BaseURL: config.ProjectManagerServiceURL,
	}
}

func (a *ProjectManagerAPI) makeAuthenticatedAPIRequest(method, path string, reqData RequestData) ([]byte, error) {
	return a.makeAuthenticatedRequest(method, fmt.Sprintf("%s/%s", a.BaseURL, path), reqData)
}
