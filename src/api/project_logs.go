package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectLogListRequest struct {
	ProjectId string `json:"projectId,omitempty"`
	Level     string `json:"level,omitempty"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
}

func (a *projectManagerAPI) ListProjectLog(filters ProjectLogListRequest) (string, error) {
	params, err := json.Marshal(filters)
	if err != nil {
		return "", fmt.Errorf("error while marshalling filters: %s", err)
	}

	resp, err := a.makeAuthenticatedAPIRequest(
		http.MethodGet,
		endpoints.log.list,
		RequestData{Params: params},
	)
	respStr := string(resp)
	if err != nil {
		fmt.Println(respStr)
		return "", err
	}

	return respStr, nil
}
