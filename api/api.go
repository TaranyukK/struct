package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	key    string
	client *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		key: apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Create(data any) error {
	return c.sendRequest("POST", "https://api.jsonbin.io/v3/b", data)
}

func (c *Client) Get(id string) error {
	return c.sendRequest("GET", fmt.Sprintf("https://api.jsonbin.io/v3/b/%s/latest", id), nil)
}

func (c *Client) Update(id string, data any) error {
	return c.sendRequest("PUT", fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id), data)
}

func (c *Client) Delete(id string) error {
	return c.sendRequest("DELETE", fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id), nil)
}

func (c *Client) List() error {
	return c.sendRequest("GET", "https://api.jsonbin.io/v3/b", nil)
}

func (c *Client) sendRequest(method, url string, data any) error {
	var bodyReader io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка маршалинга: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}

	req.Header.Set("X-Master-Key", c.key)
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("ошибка API (код %d): %s", resp.StatusCode, string(respBody))
	}

	fmt.Println(string(respBody))
	return nil
}