package api

import (
	"encoding/json"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

func (a *authAPI) Register(email, password string) (refreshToken, accessToken string, err error) {
	reqBody, err := json.Marshal(AuthRequest{Email: email, Password: password})
	if err != nil {
		return "", "", err
	}

	resp, err := a.makeAPIRequest(http.MethodPost, endpoints.auth.register, reqBody)
	if err != nil {
		return "", "", err
	}

	var authResp AuthResponse
	err = json.Unmarshal(resp, &authResp)
	if err != nil {
		return "", "", err
	}

	return authResp.RefreshToken, authResp.AccessToken, nil
}

func (a *authAPI) Login(email, password string) (refreshToken, accessToken string, err error) {
	reqBody, err := json.Marshal(AuthRequest{Email: email, Password: password})
	if err != nil {
		return "", "", err
	}

	resp, err := a.makeAPIRequest(http.MethodPost, endpoints.auth.login, reqBody)
	if err != nil {
		return "", "", err
	}

	var authResp AuthResponse
	err = json.Unmarshal(resp, &authResp)
	if err != nil {
		return "", "", err
	}

	return authResp.RefreshToken, authResp.AccessToken, nil
}

func (a *authAPI) Logout() error {
	_, err := a.makeAuthenticatedAPIRequest(http.MethodPost, endpoints.auth.logout, nil)
	return err
}
