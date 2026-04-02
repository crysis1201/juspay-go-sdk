package juspay

import (
	"context"
	"encoding/json"
	"strings"
)

// CreateRefund initiates a refund for an order.
// POST /orders/{order_id}/refunds
func (c *Client) CreateRefund(ctx context.Context, auth MerchantAuth, orderID string, req CreateRefundRequest) (*RefundResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	respBody, err := c.doRequest(ctx, "POST", "/orders/"+orderID+"/refunds", strings.NewReader(string(body)), auth)
	if err != nil {
		return nil, err
	}

	var resp RefundResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
