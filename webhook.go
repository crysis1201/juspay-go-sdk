package juspay

import "encoding/json"

// Webhook event name constants.
const (
	EventOrderSucceeded       = "ORDER_SUCCEEDED"
	EventOrderAuthorized      = "ORDER_AUTHORIZED"
	EventOrderPartialPayment  = "ORDER_PARTIAL_CHARGED"
	EventOrderVoided          = "ORDER_VOIDED"
	EventOrderFailed          = "ORDER_FAILED"
	EventOrderRefunded        = "ORDER_REFUNDED"
	EventOrderRefundFailed    = "ORDER_REFUND_FAILED"
	EventOrderRefundSucceeded = "ORDER_REFUND_SUCCEEDED"
)

// Refund status constants (value of OrderRefund.Status).
const (
	RefundStatusSuccess = "SUCCESS"
	RefundStatusFailure = "FAILURE"
	RefundStatusPending = "PENDING"
)

// ParseWebhookEvent parses a raw Juspay webhook payload into a WebhookEvent.
func ParseWebhookEvent(payload []byte) (*WebhookEvent, error) {
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, err
	}
	return &event, nil
}
