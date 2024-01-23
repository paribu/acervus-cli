package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/paribu/acervus-cli/src/settings"
)

type TestRequest struct {
	Yaml    string `json:"yaml"`
	Abi     string `json:"abi"`
	Schema  string `json:"schema"`
	Project string `json:"project"`
}

type TestResponse struct {
	Results []Result `json:"results"`
	Data    []Data   `json:"data"`
}

type Result struct {
	Event        Event  `json:"event"`
	Status       bool   `json:"status"`
	Data         []Data `json:"data"`
	Logs         []Log  `json:"logs"`
	ErrorMessage string `json:"errorMessage"`
}

type Data struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
}

type Event struct {
	ContractAddress string   `json:"contractAddress"`
	BlockNumber     uint32   `json:"blockNumber"`
	BlockHash       string   `json:"blockHash"`
	TransactionHash string   `json:"transactionHash"`
	LogIndex        uint32   `json:"logIndex"`
	EventCount      uint32   `json:"eventCount"`
	Topics          []string `json:"topics"`
	Data            string   `json:"data"`
	IsCRUDMode      bool     `json:"isCrudMode"`
}

type Log struct {
	Log       string `json:"log"`
	Level     string `json:"level"`
	CreatedAt string `json:"createdAt"`
}

func (a *projectManagerAPI) Test(projectID, settingsFilePath, projectFilePath string) (*TestResponse, error) {
	settingsFile, err := settings.NewProjectFromFile(settingsFilePath)
	if err != nil {
		return nil, err
	}

	settingsStr, err := settingsFile.ToString()
	if err != nil {
		return nil, err
	}

	if projectFilePath == "" {
		projectFilePath = fmt.Sprintf("./project/%s/project.ts", projectID)
	}

	projectStr, err := os.ReadFile(projectFilePath)
	if err != nil {
		return nil, err
	}

	abiStr, err := os.ReadFile(settingsFile.Sources[0].Source.Abi)
	if err != nil {
		return nil, err
	}

	schemaStr, err := os.ReadFile(settingsFile.Schema)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(DeployRequest{
		Yaml:    settingsStr,
		Abi:     string(abiStr),
		Schema:  string(schemaStr),
		Project: string(projectStr),
	})
	if err != nil {
		return nil, err
	}

	resp, err := a.makeAuthenticatedAPIRequest(
		http.MethodPost,
		endpoints.project.test(projectID),
		RequestData{Body: body},
	)
	if err != nil {
		return nil, err
	}

	var testResp TestResponse
	err = json.Unmarshal(resp, &testResp)
	if err != nil {
		return nil, errors.New(string(resp))
	}

	return &testResp, nil
}
