package juspay

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is a stateless HTTP client for the Juspay API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// Config holds the configuration for creating a Juspay Client.
type Config struct {
	BaseURL    string       // e.g. "https://api.juspay.in"
	HTTPClient *http.Client // optional; defaults to a client with 30s timeout
}

// MerchantAuth holds the credentials for authenticating API requests to a specific merchant.
type MerchantAuth struct {
	APIKey     string
	MerchantID string
}

// NewClient creates a new Juspay API client.
func NewClient(cfg Config) *Client {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	}
	return &Client{
		baseURL:    cfg.BaseURL,
		httpClient: cfg.HTTPClient,
	}
}

// doRequest is the internal helper for making authenticated HTTP requests.
func (c *Client) doRequest(ctx context.Context, method, path string, body io.Reader, auth MerchantAuth) ([]byte, error) {
	reqURL := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(auth.APIKey, "")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-merchantid", "zuzu")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Body:       string(respBody),
		}
	}

	return respBody, nil
}

// APIError represents a non-2xx response from the Juspay API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("juspay API returned status %d: %s", e.StatusCode, e.Body)
}
