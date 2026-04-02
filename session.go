package juspay

import (
	"context"
	"encoding/json"
	"strings"
)

// CreateSession creates a Juspay HyperCheckout payment session.
// POST /session
func (c *Client) CreateSession(ctx context.Context, auth MerchantAuth, req CreateSessionRequest) (*SessionResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	respBody, err := c.doRequest(ctx, "POST", "/session", strings.NewReader(string(body)), auth)
	if err != nil {
		return nil, err
	}

	var resp SessionResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
