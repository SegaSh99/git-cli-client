package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git-cli-client/internal/utils"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) Login(username, password string) (string, error) {
	authRequest := AuthRequest{
		Username: username,
		Password: password,
	}

	requestBody, err := json.Marshal(authRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v1/sign_in", c.BaseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer utils.CloseResponseBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body := utils.ReadResponseBody(resp)
		return "", fmt.Errorf("failed to login, status code: %d\nMessage:%s", resp.StatusCode, body)
	}

	var authResponse utils.AuthToken
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return "", err
	}

	if err := utils.SaveToken(authResponse); err != nil {
		return "", err
	}

	return authResponse.AccessToken, nil
}
