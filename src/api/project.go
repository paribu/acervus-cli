package api

import (
	"encoding/json"
	"net/http"

	"github.com/paribu/acervus-cli/src/settings"
)

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Network     string `json:"network"`
}

type CreateProjectResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Network     string `json:"network"`
	UserId      string `json:"userId"`
	ProjectId   string `json:"projectId"`
}

type PaginationResult struct {
	Results   []ProjectItem `json:"results"`
	PageTotal int           `json:"pageTotal"`
	Total     int           `json:"total"`
}

type ProjectItem struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Abi         string `json:"abi"`
	Yaml        string `json:"yaml"`
	Schema      string `json:"schema"`
	Address     string `json:"address"`
	Topic       string `json:"topic"`
	StartBlock  int64  `json:"startBlock"`
	EndBlock    int64  `json:"endBlock"`
	IsDeleted   bool   `json:"isDeleted"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func (a *projectManagerAPI) CreateProject(settingsFilepath string) (*CreateProjectResponse, error) {
	yamlFile, err := settings.NewProjectFromFile(settingsFilepath)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(CreateProjectRequest{
		Name:        yamlFile.Project,
		Description: yamlFile.Description,
		Network:     yamlFile.Sources[0].Network,
	})
	if err != nil {
		return nil, err
	}

	resp, err := a.makeAuthenticatedAPIRequest(http.MethodPost, endpoints.project.create, body)
	if err != nil {
		return nil, err
	}

	var createProjectResp CreateProjectResponse
	err = json.Unmarshal(resp, &createProjectResp)
	if err != nil {
		return nil, err
	}

	return &createProjectResp, nil
}

func (a *projectManagerAPI) ListProjects() ([]ProjectItem, error) {
	resp, err := a.makeAuthenticatedAPIRequest(http.MethodGet, endpoints.project.list+"?page=1&limit=0", nil)
	if err != nil {
		return []ProjectItem{}, err
	}

	var listResponse PaginationResult
	err = json.Unmarshal(resp, &listResponse)
	if err != nil {
		return []ProjectItem{}, err
	}

	return listResponse.Results, nil
}

func (a *projectManagerAPI) DeleteProject(projectID string) error {
	_, err := a.makeAuthenticatedAPIRequest(http.MethodDelete, endpoints.project.delete(projectID), nil)
	return err
}

func (a *projectManagerAPI) ExportProject(projectID string) error {
	_, err := a.makeAuthenticatedAPIRequest(http.MethodPost, endpoints.project.export(projectID), nil)
	return err
}
