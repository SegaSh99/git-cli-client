package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git-cli-client/internal/utils"
	"git-cli-client/pkg/models"
	"net/http"
)

func (c *Client) CreateUser(user models.CreateUser) error {
	requestBody, err := json.Marshal(user)
	if err != nil {
		return err
	}

	resp, err := c.DoRequest(http.MethodPost, "/v1/admin/users", bytes.NewReader(requestBody))
	if err != nil {
		return err
	}
	defer utils.CloseResponseBody(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		body := utils.ReadResponseBody(resp)
		return fmt.Errorf("failed to create user, status code: %d\nMessage:%s", resp.StatusCode, body)
	}

	return nil
}

func (c *Client) GetUsers() ([]models.User, error) {
	resp, err := c.DoRequest(http.MethodGet, "/v1/admin/users", nil)
	if err != nil {
		return nil, err
	}
	defer utils.CloseResponseBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body := utils.ReadResponseBody(resp)
		return nil, fmt.Errorf("failed to get users, status code: %d\nMessage:%s", resp.StatusCode, body)
	}

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) UpdateUser(u models.UpdateUser) error {
	requestBody, err := json.Marshal(u)
	if err != nil {
		return err
	}

	resp, err := c.DoRequest(http.MethodPatch, "/v1/admin/users/"+u.Username, bytes.NewReader(requestBody))
	if err != nil {
		return err
	}
	defer utils.CloseResponseBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body := utils.ReadResponseBody(resp)
		return fmt.Errorf("failed to activate user, status code: %d\nMessage:%s", resp.StatusCode, body)
	}

	return nil
}
