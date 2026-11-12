package juspay

import (
	"context"
	"encoding/json"
	"net/url"
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

// CaptureTransaction captures a pre-authorized (AUTHORIZED) transaction. For a partial capture
// set req.Amount; Juspay voids the remaining authorization. The response is usually
// CAPTURE_INITIATED, with the final status delivered via webhook / order status.
// POST /v2/txns/{txn_uuid}/capture
func (c *Client) CaptureTransaction(ctx context.Context, auth MerchantAuth, txnUUID string, req CaptureTransactionRequest) (*TxnModificationResponse, []byte, error) {
	form := url.Values{}
	if req.Amount != "" {
		form.Set("amount", req.Amount)
	}
	if req.IdempotenceKey != "" {
		form.Set("idempotence_key", req.IdempotenceKey)
	}

	respBody, err := c.doRequestWithContentType(ctx, "POST", "/v2/txns/"+url.PathEscape(txnUUID)+"/capture",
		"application/x-www-form-urlencoded", strings.NewReader(form.Encode()), auth)
	if err != nil {
		return nil, nil, err
	}

	var resp TxnModificationResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, nil, err
	}
	return &resp, respBody, nil
}

// VoidTransaction releases a pre-authorized transaction. The response is usually VOID_INITIATED,
// with the final status (VOIDED / VOID_FAILED) delivered via webhook / order status.
// POST /v2/txns/{txn_uuid}/void
func (c *Client) VoidTransaction(ctx context.Context, auth MerchantAuth, txnUUID string) (*TxnModificationResponse, []byte, error) {
	respBody, err := c.doRequest(ctx, "POST", "/v2/txns/"+url.PathEscape(txnUUID)+"/void", nil, auth)
	if err != nil {
		return nil, nil, err
	}

	var resp TxnModificationResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, nil, err
	}
	return &resp, respBody, nil
}
