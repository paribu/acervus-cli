package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectDataListRequest struct {
	ProjectId string `json:"projectId,omitempty"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"valueFilters,omitempty"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
}

type ProjectDataListResponse struct {
	Results   []ProjectDataItem `json:"results"`
	PageTotal int               `json:"pageTotal"`
	Total     int               `json:"total"`
}

type ProjectDataItem struct {
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

func (a *projectManagerAPI) ListProjectData(filters ProjectDataListRequest) (string, error) {
	params, err := json.Marshal(filters)
	if err != nil {
		return "", fmt.Errorf("error while marshalling filters: %s", err)
	}

	resp, err := a.makeAuthenticatedAPIRequest(
		http.MethodGet,
		endpoints.data.list,
		RequestData{Params: params},
	)
	respStr := string(resp)
	if err != nil {
		fmt.Println(respStr)
		return "", err
	}

	return respStr, nil
}
