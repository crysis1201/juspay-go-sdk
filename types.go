package juspay

// ──────────────────────────────────────────────────────────────────────────────
// Session (POST /session)
// ──────────────────────────────────────────────────────────────────────────────

type CreateSessionRequest struct {
	OrderID             string        `json:"order_id"`
	Amount              string        `json:"amount"`
	CustomerID          string        `json:"customer_id,omitempty"`
	CustomerEmail       *string       `json:"customer_email,omitempty"`
	CustomerPhone       *string       `json:"customer_phone,omitempty"`
	PaymentPageClientID string        `json:"payment_page_client_id"`
	Action              string        `json:"action"`
	ReturnURL           *string       `json:"return_url,omitempty"`
	Currency            string        `json:"currency"`
	Description         string        `json:"description,omitempty"`
	FirstName           string        `json:"first_name,omitempty"`
	LastName            string        `json:"last_name,omitempty"`
	UDF1                string        `json:"udf1,omitempty"`
	PaymentRules        *PaymentRules `json:"payment_rules,omitempty"`
}

// PaymentRules is the top-level Juspay payment_rules object.
type PaymentRules struct {
	PaymentFlows *PaymentFlows `json:"payment_flows,omitempty"`
}

// PaymentFlows contains the payment instrument rules.
type PaymentFlows struct {
	PaymentInstrumentRules *PaymentInstrumentRules `json:"payment_instrument_rules,omitempty"`
}

// PaymentInstrumentRules defines the status and variant info.
type PaymentInstrumentRules struct {
	Status string                      `json:"status"`
	Info   *PaymentInstrumentRulesInfo `json:"info,omitempty"`
}

// PaymentInstrumentRulesInfo holds the list of payment variants.
type PaymentInstrumentRulesInfo struct {
	Variants []PaymentVariant `json:"variants"`
}

// PaymentVariant represents a single payment method variant with amount and filter.
type PaymentVariant struct {
	Amount        VariantAmount  `json:"amount"`
	OverrideRules *OverrideRules `json:"override_rules,omitempty"`
}

// VariantAmount defines the amount type and value for a variant.
type VariantAmount struct {
	AmountType string `json:"amount_type"`
	Value      string `json:"value"`
}

// OverrideRules contains the payment filter for a variant.
type OverrideRules struct {
	PaymentFilter *PaymentFilter `json:"payment_filter,omitempty"`
}

// PaymentFilter controls which payment methods are allowed.
type PaymentFilter struct {
	AllowDefaultOptions bool                  `json:"allowDefaultOptions"`
	Options             []PaymentFilterOption `json:"options"`
}

// PaymentFilterOption enables or disables a specific payment method type.
type PaymentFilterOption struct {
	PaymentMethodType string `json:"paymentMethodType"`
	Enable            bool   `json:"enable"`
}

type SessionResponse struct {
	OrderID      string        `json:"order_id"`
	Status       string        `json:"status"`
	PaymentLinks *PaymentLinks `json:"payment_links,omitempty"`
}

type PaymentLinks struct {
	Web    *string `json:"web,omitempty"`
	Mobile *string `json:"mobile,omitempty"`
	Iframe *string `json:"iframe,omitempty"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Cards (GET /cards)
// ──────────────────────────────────────────────────────────────────────────────

type Card struct {
	CardToken      string `json:"card_token"`
	CardReference  string `json:"card_reference"`
	CardNumber     string `json:"card_number"`
	CardBrand      string `json:"card_brand"`
	CardType       string `json:"card_type"`
	ExpiryMonth    string `json:"expiry_month"`
	ExpiryYear     string `json:"expiry_year"`
	LastFourDigits string `json:"last_four_digits"`
	SavedToLocker  bool   `json:"saved_to_locker"`
}

type CardsResponse struct {
	Cards []Card `json:"cards"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Transaction (POST /txns)
// ──────────────────────────────────────────────────────────────────────────────

type CreateTransactionRequest struct {
	MerchantID           string            `json:"merchant_id"`
	PaymentMethodType    string            `json:"payment_method_type,omitempty"`
	CardToken            string            `json:"card_token"`
	SaveToLocker         bool              `json:"save_to_locker"`
	RedirectAfterPayment bool              `json:"redirect_after_payment"`
	Format               string            `json:"format"`
	AuthType             string            `json:"auth_type"`
	Order                TxnOrder          `json:"order"`
	Metadata             map[string]any    `json:"metadata,omitempty"`
}

type TxnOrder struct {
	OrderID    string `json:"order_id"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	CustomerID string `json:"customer_id"`
	UDF1       string `json:"udf1,omitempty"`
}

type TransactionResponse struct {
	TxnID  string `json:"txn_id"`
	Status string `json:"status"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Refund (POST /orders/{order_id}/refunds)
// ──────────────────────────────────────────────────────────────────────────────

type CreateRefundRequest struct {
	OrderID         string `json:"order_id"`
	Amount          string `json:"amount"`
	UniqueRequestID string `json:"unique_request_id"`
	UDF1            string `json:"udf1,omitempty"`
}

type RefundResponse struct {
	ID     string  `json:"id"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Order (POST /orders)
// ──────────────────────────────────────────────────────────────────────────────

type CreateOrderRequest struct {
	OrderID    string `json:"order_id"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	CustomerID string `json:"customer_id,omitempty"`
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Order Status (GET /orders/{order_id})
// ──────────────────────────────────────────────────────────────────────────────

type OrderStatusResponse struct {
	OrderID           string        `json:"order_id"`
	MerchantID        string        `json:"merchant_id"`
	Status            string        `json:"status"`
	StatusID          int           `json:"status_id"`
	Amount            float64       `json:"amount"`
	Currency          string        `json:"currency"`
	TxnID             string        `json:"txn_id"`
	PaymentMethodType string        `json:"payment_method_type"`
	Card              *Card         `json:"card,omitempty"`
	Refunds           []OrderRefund `json:"refunds,omitempty"`
}

type OrderRefund struct {
	UniqueRequestID string  `json:"unique_request_id"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Webhook
// ──────────────────────────────────────────────────────────────────────────────

type WebhookEvent struct {
	EventName   string         `json:"event_name"`
	DateCreated string         `json:"date_created"`
	Content     WebhookContent `json:"content"`
}

type WebhookContent struct {
	Order WebhookOrder `json:"order"`
}

type WebhookOrder struct {
	OrderID           string        `json:"order_id"`
	MerchantID        string        `json:"merchant_id"`
	Status            string        `json:"status"`
	StatusID          int           `json:"status_id"`
	Amount            float64       `json:"amount"`
	Currency          string        `json:"currency"`
	TxnID             string        `json:"txn_id"`
	PaymentMethodType string        `json:"payment_method_type"`
	Card              *Card         `json:"card,omitempty"`
	Refunds           []OrderRefund `json:"refunds,omitempty"`
	UDF1              string        `json:"udf1,omitempty"`
}
