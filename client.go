package tinkoff

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) do(url, token string, unmarshal interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("build new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("provider do: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(unmarshal)
	if err != nil {
		return fmt.Errorf("decode json: %w", err)
	}

	return nil
}
