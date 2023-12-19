package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/paribu/acervus-cli/src/config"
	"github.com/paribu/acervus-cli/src/credential"
)

const NO_ACCESS_TOKEN, NO_REFRESH_TOKEN = "", ""
const CLI_STATUS_CODE_INTERNAL_ERROR = -1

type apiClient struct{}

type RequestData struct {
	Body   []byte `json:"body"`
	Params []byte `json:"params"`
}

func (a *apiClient) makeAuthenticatedRequest(method, url string, reqData RequestData) ([]byte, error) {
	cred, err := credential.GetCurrentCredential()
	if err != nil {
		return []byte{}, err
	}

	respBody, respStatusCode, err := a.doRequest(method, url, cred.AccessToken, reqData)
	if err != nil {
		return []byte{}, err
	}

	if respStatusCode == http.StatusUnauthorized {
		err := a.refreshAccessToken()
		if err != nil {
			return []byte{}, err
		}

		return a.makeAuthenticatedRequest(method, url, reqData)
	}

	if !isResponseOk(respStatusCode) {
		return respBody, formatRespError(respStatusCode, respBody)
	}

	return respBody, nil
}

func (a *apiClient) makeRequest(method, url string, reqData RequestData) ([]byte, error) {
	respBody, statusCode, err := a.doRequest(method, url, NO_ACCESS_TOKEN, reqData)
	if err != nil {
		return []byte{}, err
	}

	if !isResponseOk(statusCode) {
		return respBody, formatRespError(statusCode, respBody)
	}

	return respBody, nil
}

func (a *apiClient) doRequest(method, url, accessToken string, reqData RequestData) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqData.Body))
	if err != nil {
		return []byte{}, CLI_STATUS_CODE_INTERNAL_ERROR, err
	}

	if reqData.Params != nil {
		// There is params provided for GET request,
		// so we need to add them to the URL

		paramsObj := make(map[string]interface{})
		err = json.Unmarshal(reqData.Params, &paramsObj)
		if err != nil {
			return []byte{}, CLI_STATUS_CODE_INTERNAL_ERROR, fmt.Errorf("failed to unmarshal params: %w", err)
		}

		q := req.URL.Query()
		for k, v := range paramsObj {
			q.Add(k, fmt.Sprintf("%v", v))
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("Content-Type", "application/json")

	if accessToken != NO_ACCESS_TOKEN {
		req.Header.Set(
			"Authorization",
			fmt.Sprintf("Bearer %s", accessToken),
		)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: config.AllowUnsignedCertificates,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, CLI_STATUS_CODE_INTERNAL_ERROR, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, resp.StatusCode, err
	}

	return respBody, resp.StatusCode, nil
}

func (a *apiClient) refreshAccessToken() error {
	cred, err := credential.GetCurrentCredential()
	if err != nil {
		return err
	}

	refreshTokenResponse, refreshTokenStatusCode, err := a.doRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", config.AuthServiceURL, endpoints.auth.refreshToken),
		cred.RefreshToken,
		RequestData{},
	)
	if err != nil {
		return err
	}

	if !isResponseOk(refreshTokenStatusCode) {
		return formatRespError(refreshTokenStatusCode, refreshTokenResponse)
	}

	var authResp AuthResponse
	err = json.Unmarshal(refreshTokenResponse, &authResp)
	if err != nil {
		return err
	}

	if authResp.RefreshToken == NO_REFRESH_TOKEN || authResp.AccessToken == NO_ACCESS_TOKEN {
		return errors.New("failed to refresh token: new refresh token or access token is empty after refresh")
	}

	err = credential.UpdateCredential(cred.Email, authResp.RefreshToken, authResp.AccessToken)
	if err != nil {
		return err
	}

	return nil
}

func isResponseOk(statusCode int) bool {
	return statusCode == http.StatusOK ||
		statusCode == http.StatusAccepted ||
		statusCode == http.StatusCreated
}

func formatRespError(statusCode int, respBody []byte) error {
	var respFields map[string]interface{}
	err := json.Unmarshal(respBody, &respFields)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	message, ok := respFields["message"].(string)

	errMsg := ""
	if ok && message != "" {
		errMsg = fmt.Sprintf("with message: %s", message)
	}

	return fmt.Errorf("server returned status %d %s", statusCode, errMsg)
}
