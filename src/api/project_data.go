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
