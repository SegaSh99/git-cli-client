package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git-cli-client/internal/utils"
	"git-cli-client/pkg/models"
	"io"
	"net/http"
)

type CollaboratorApi struct {
	Owner      string
	Repo       string
	User       string
	Permission string
}

func (c *Client) GetRepositories() ([]models.Repository, error) {
	resp, err := c.DoRequest(http.MethodGet, "/v1/repos/search", nil)
	if err != nil {
		return nil, err
	}
	defer utils.CloseResponseBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body := utils.ReadResponseBody(resp)
		return nil, fmt.Errorf("failed to get repositories, status code: %d\nMessage:%s", resp.StatusCode, body)
	}

	var repos []models.Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func (c *Client) GetRepoMembers(ca CollaboratorApi) ([]models.User, error) {
	url := fmt.Sprintf("/v1/repos/%s/%s/collaborators", ca.Owner, ca.Repo)
	resp, err := c.DoRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer utils.CloseResponseBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body := utils.ReadResponseBody(resp)
		return nil, fmt.Errorf("failed to get repository members, status code: %d\nMessage:%s", resp.StatusCode, body)
	}

	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) AddRepoMember(ca CollaboratorApi) error {
	url := fmt.Sprintf("/v1/repos/%s/%s/collaborators/%s", ca.Owner, ca.Repo, ca.User)

	var requestBody io.Reader
	if ca.Permission != "" {
		body, err := json.Marshal(map[string]string{"permission": ca.Permission})
		if err != nil {
			return err
		}
		requestBody = bytes.NewBuffer(body)
	} else {
		requestBody = nil
	}

	resp, err := c.DoRequest(http.MethodPut, url, requestBody)
	if err != nil {
		return err
	}
	defer utils.CloseResponseBody(resp.Body)

	if resp.StatusCode != http.StatusNoContent {
		body := utils.ReadResponseBody(resp)
		return fmt.Errorf("failed to add repository member, status code: %d\nMessage:%s", resp.StatusCode, body)
	}

	return nil
}

func (c *Client) DeleteRepoMember(ca CollaboratorApi) error {
	url := fmt.Sprintf("/v1/repos/%s/%s/collaborators/%s", ca.Owner, ca.Repo, ca.User)
	resp, err := c.DoRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	defer utils.CloseResponseBody(resp.Body)

	if resp.StatusCode != http.StatusNoContent {
		body := utils.ReadResponseBody(resp)
		return fmt.Errorf("failed to add repository member, status code: %d\nMessage:%s", resp.StatusCode, body)
	}

	return nil
}
