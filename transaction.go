package juspay

import (
	"context"
	"encoding/json"
	"strings"
)

// CreateTransaction initiates a transaction (e.g. charge a saved card).
// POST /txns
func (c *Client) CreateTransaction(ctx context.Context, auth MerchantAuth, req CreateTransactionRequest) (*TransactionResponse, []byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	respBody, err := c.doRequest(ctx, "POST", "/txns", strings.NewReader(string(body)), auth)
	if err != nil {
		return nil, nil, err
	}

	var resp TransactionResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, nil, err
	}
	return &resp, respBody, nil
}
