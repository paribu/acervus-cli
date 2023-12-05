package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/paribu/acervus-cli/src/settings"
)

type DeployRequest struct {
	Yaml    string `json:"yaml"`
	Abi     string `json:"abi"`
	Schema  string `json:"schema"`
	Project string `json:"project"`
}

type DeployResponse struct {
	ResultMessage string `json:"resultMessage"`
}

func (a *projectManagerAPI) Deploy(projectID, settingsFilePath, projectFilePath string) (string, error) {
	settingsFile, err := settings.NewProjectFromFile(settingsFilePath)
	if err != nil {
		return "", err
	}

	settingsStr, err := settingsFile.ToString()
	if err != nil {
		return "", err
	}

	if projectFilePath == "" {
		projectFilePath = fmt.Sprintf("./project/%s/project.ts", projectID)
	}

	projectStr, err := os.ReadFile(projectFilePath)
	if err != nil {
		return "", err
	}

	abiStr, err := os.ReadFile(settingsFile.Sources[0].Source.Abi)
	if err != nil {
		return "", err
	}

	schemaStr, err := os.ReadFile(settingsFile.Schema)
	if err != nil {
		return "", err
	}

	body, err := json.Marshal(DeployRequest{
		Yaml:    settingsStr,
		Abi:     string(abiStr),
		Schema:  string(schemaStr),
		Project: string(projectStr),
	})
	if err != nil {
		return "", err
	}

	_, err = a.makeAuthenticatedAPIRequest(http.MethodPost, endpoints.project.deploy(projectID), body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Deployed: %s", projectID), nil
}
