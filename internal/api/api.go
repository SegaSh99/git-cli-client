package api

import (
	"io"
	"net/http"
	"os"
)

type Client struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		BaseURL: os.Getenv("BASE_URL"),
		Token:   token,
		Client:  &http.Client{},
	}
}

func (c *Client) DoRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.BaseURL+endpoint, body)
	if err != nil {
		return nil, err
	}

	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
	req.Header.Set("Content-Type", "application/json")

	return c.Client.Do(req)
}
