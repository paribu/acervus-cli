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

func (a *apiClient) makeAuthenticatedRequest(method, url string, body []byte) ([]byte, error) {
	cred, err := credential.GetCurrentCredential()
	if err != nil {
		return []byte{}, err
	}

	body, statusCode, err := a.doRequest(method, url, cred.AccessToken, body)
	if err != nil {
		return []byte{}, err
	}

	if statusCode == http.StatusUnauthorized {
		err := a.refreshAccessToken()
		if err != nil {
			return []byte{}, err
		}

		return a.makeAuthenticatedRequest(method, url, body)
	}

	if !isResponseOk(statusCode) {
		return body, formatError(statusCode, body)
	}

	return body, nil
}

func (a *apiClient) makeRequest(method, url string, body []byte) ([]byte, error) {
	body, statusCode, err := a.doRequest(method, url, NO_ACCESS_TOKEN, body)
	if err != nil {
		return []byte{}, err
	}

	if !isResponseOk(statusCode) {
		return body, formatError(statusCode, body)
	}

	return body, nil
}

func (a *apiClient) doRequest(method, url, accessToken string, body []byte) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, CLI_STATUS_CODE_INTERNAL_ERROR, err
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
		nil,
	)
	if err != nil {
		return err
	}

	if !isResponseOk(refreshTokenStatusCode) {
		return formatError(refreshTokenStatusCode, refreshTokenResponse)
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

func formatError(statusCode int, body []byte) error {
	var respFields map[string]interface{}
	json.Unmarshal(body, &respFields)

	message, ok := respFields["message"].(string)

	errMsg := ""
	if ok && message != "" {
		errMsg = fmt.Sprintf("with message: %s", message)
	}

	return fmt.Errorf("server returned status %d %s", statusCode, errMsg)
}
