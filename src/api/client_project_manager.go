package api

import "github.com/paribu/acervus-cli/src/config"

type projectManagerAPI struct {
	apiClient
	BaseURL string
}

func NewProjectManagerAPI() *projectManagerAPI {
	return &projectManagerAPI{
		BaseURL: config.ProjectManagerServiceURL,
	}
}
