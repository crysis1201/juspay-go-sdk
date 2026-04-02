package juspay

import (
	"context"
	"encoding/json"
	"strings"
)

// CreateOrder creates a new order.
// POST /orders
func (c *Client) CreateOrder(ctx context.Context, auth MerchantAuth, req CreateOrderRequest) (*OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	respBody, err := c.doRequest(ctx, "POST", "/orders", strings.NewReader(string(body)), auth)
	if err != nil {
		return nil, err
	}

	var resp OrderResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrderStatus fetches the current status of an order.
// GET /orders/{order_id}
func (c *Client) GetOrderStatus(ctx context.Context, auth MerchantAuth, orderID string) (*OrderStatusResponse, error) {
	respBody, err := c.doRequest(ctx, "GET", "/orders/"+orderID, nil, auth)
	if err != nil {
		return nil, err
	}

	var resp OrderStatusResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
