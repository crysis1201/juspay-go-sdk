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
	CardToken               string         `json:"card_token,omitempty"`
	CardReference           string         `json:"card_reference,omitempty"`
	CardNumber              string         `json:"card_number,omitempty"`
	CardBrand               string         `json:"card_brand,omitempty"`
	CardType                string         `json:"card_type,omitempty"`
	CardSubType             *string        `json:"card_sub_type,omitempty"`
	CardSubTypeCategory     any            `json:"card_sub_type_category,omitempty"`
	ExtendedCardType        string         `json:"extended_card_type,omitempty"`
	CardIssuer              string         `json:"card_issuer,omitempty"`
	CardIssuerCountry       string         `json:"card_issuer_country,omitempty"`
	CardIsin                string         `json:"card_isin,omitempty"`
	CardFingerprint         string         `json:"card_fingerprint,omitempty"`
	CardExpMonth            string         `json:"card_exp_month,omitempty"`
	CardExpYear             string         `json:"card_exp_year,omitempty"`
	ExpiryMonth             string         `json:"expiry_month,omitempty"`
	ExpiryYear              string         `json:"expiry_year,omitempty"`
	LastFourDigits          string         `json:"last_four_digits,omitempty"`
	NameOnCard              string         `json:"name_on_card,omitempty"`
	Nickname                string         `json:"nickname,omitempty"`
	CountryCode             string         `json:"country_code,omitempty"`
	JuspayBankCode          any            `json:"juspay_bank_code,omitempty"`
	Provider                string         `json:"provider,omitempty"`
	ProviderCategory        string         `json:"provider_category,omitempty"`
	VaultProvider           string         `json:"vault_provider,omitempty"`
	Expired                 bool           `json:"expired,omitempty"`
	TokenizeSupport         bool           `json:"tokenize_support,omitempty"`
	SavedToLocker           bool           `json:"saved_to_locker,omitempty"`
	UsingSavedCard          bool           `json:"using_saved_card,omitempty"`
	UsingToken              bool           `json:"using_token,omitempty"`
	TokenType               string         `json:"token_type,omitempty"`
	Tokens                  []any          `json:"tokens,omitempty"`
	PaymentAccountReference string         `json:"payment_account_reference,omitempty"`
	Metadata                map[string]any `json:"metadata,omitempty"`
}

type CardsResponse struct {
	Cards      []Card `json:"cards"`
	MerchantID string `json:"merchantId,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Transaction (POST /txns)
// ──────────────────────────────────────────────────────────────────────────────

type CreateTransactionRequest struct {
	MerchantID           string `json:"merchant_id"`
	CardToken            string `json:"card_token"`
	SaveToLocker         bool   `json:"save_to_locker"`
	RedirectAfterPayment bool   `json:"redirect_after_payment"`
	Format               string `json:"format"`
	AuthType             string `json:"auth_type"`
	OrderID              string `json:"order.order_id"`
	Amount               string `json:"order.amount"`
	Currency             string `json:"order.currency"`
	CustomerID           string `json:"order.customer_id"`
	UDF1                 string `json:"order.udf1,omitempty"`
	AllowCardNo3DS       string `json:"order.metadata.txns.allow_card_no_3ds,omitempty"`
}

type TransactionResponse struct {
	OrderID      string           `json:"order_id,omitempty"`
	TxnID        string           `json:"txn_id,omitempty"`
	TxnUUID      string           `json:"txn_uuid,omitempty"`
	Status       string           `json:"status,omitempty"`
	Payment      *TxnPayment      `json:"payment,omitempty"`
	OfferDetails *TxnOfferDetails `json:"offer_details,omitempty"`
}

type TxnPayment struct {
	Authentication *TxnAuthentication `json:"authentication,omitempty"`
}

type TxnAuthentication struct {
	Method string `json:"method,omitempty"`
	URL    string `json:"url,omitempty"`
}

type TxnOfferDetails struct {
	Offers []any `json:"offers,omitempty"`
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
	ID                          string         `json:"id,omitempty"`
	OrderID                     string         `json:"order_id,omitempty"`
	MerchantID                  string         `json:"merchant_id,omitempty"`
	Status                      string         `json:"status,omitempty"`
	StatusID                    int            `json:"status_id,omitempty"`
	Amount                      float64        `json:"amount,omitempty"`
	EffectiveAmount             float64        `json:"effective_amount,omitempty"`
	AmountRefunded              float64        `json:"amount_refunded,omitempty"`
	MaximumEligibleRefundAmount float64        `json:"maximum_eligible_refund_amount,omitempty"`
	Currency                    string         `json:"currency,omitempty"`
	TxnID                       string         `json:"txn_id,omitempty"`
	TxnUUID                     string         `json:"txn_uuid,omitempty"`
	GatewayID                   int            `json:"gateway_id,omitempty"`
	PaymentMethod               string         `json:"payment_method,omitempty"`
	PaymentMethodType           string         `json:"payment_method_type,omitempty"`
	AuthType                    string         `json:"auth_type,omitempty"`
	Refunded                    bool           `json:"refunded"`
	DateCreated                 string         `json:"date_created,omitempty"`
	ReturnURL                   string         `json:"return_url,omitempty"`
	Description                 string         `json:"description,omitempty"`
	ProductID                   string         `json:"product_id,omitempty"`
	CustomerID                  string         `json:"customer_id,omitempty"`
	CustomerEmail               string         `json:"customer_email,omitempty"`
	CustomerPhone               string         `json:"customer_phone,omitempty"`
	BankPG                      any            `json:"bank_pg,omitempty"`
	BankErrorCode               string         `json:"bank_error_code,omitempty"`
	BankErrorMessage            string         `json:"bank_error_message,omitempty"`
	GatewayReferenceID          any            `json:"gateway_reference_id,omitempty"`
	RewardsBreakup              any            `json:"rewards_breakup,omitempty"`
	Card                        *Card          `json:"card,omitempty"`
	TxnDetail                   *TxnDetail     `json:"txn_detail,omitempty"`
	Refunds                     []OrderRefund  `json:"refunds,omitempty"`
	Offers                      []any          `json:"offers,omitempty"`
	PaymentLinks                *PaymentLinks  `json:"payment_links,omitempty"`
	Metadata                    map[string]any `json:"metadata,omitempty"`
	PaymentGatewayResponse      map[string]any `json:"payment_gateway_response,omitempty"`
	UDF1                        string         `json:"udf1,omitempty"`
	UDF2                        string         `json:"udf2,omitempty"`
	UDF3                        string         `json:"udf3,omitempty"`
	UDF4                        string         `json:"udf4,omitempty"`
	UDF5                        string         `json:"udf5,omitempty"`
	UDF6                        string         `json:"udf6,omitempty"`
	UDF7                        string         `json:"udf7,omitempty"`
	UDF8                        string         `json:"udf8,omitempty"`
	UDF9                        string         `json:"udf9,omitempty"`
	UDF10                       string         `json:"udf10,omitempty"`
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
	// identity
	ID         string `json:"id,omitempty"`
	OrderID    string `json:"order_id,omitempty"`
	MerchantID string `json:"merchant_id,omitempty"`

	// customer
	CustomerEmail            string `json:"customer_email,omitempty"`
	CustomerPhone            string `json:"customer_phone,omitempty"`
	CustomerID               string `json:"customer_id,omitempty"`
	CustomerPhoneCountryCode any    `json:"customer_phone_country_code,omitempty"`

	// status & amount
	Status                      string  `json:"status,omitempty"`
	StatusID                    int     `json:"status_id,omitempty"`
	Amount                      float64 `json:"amount,omitempty"`
	EffectiveAmount             float64 `json:"effective_amount,omitempty"`
	AmountRefunded              float64 `json:"amount_refunded,omitempty"`
	PaidAmount                  any     `json:"paid_amount,omitempty"`
	MaximumEligibleRefundAmount float64 `json:"maximum_eligible_refund_amount,omitempty"`
	Currency                    string  `json:"currency,omitempty"`
	Refunded                    bool    `json:"refunded,omitempty"`

	// timestamps
	DateCreated string `json:"date_created,omitempty"`
	LastUpdated string `json:"last_updated,omitempty"`

	// misc
	ReturnURL   string   `json:"return_url,omitempty"`
	ProductID   string   `json:"product_id,omitempty"`
	Description string   `json:"description,omitempty"`
	Conflicted  bool     `json:"conflicted,omitempty"`
	NextAction  []string `json:"next_action,omitempty"`

	// CHARGED-only top-level (for PARTIAL_CHARGED these live in TxnList)
	TxnID                  string         `json:"txn_id,omitempty"`
	TxnUUID                string         `json:"txn_uuid,omitempty"`
	PaymentMethodType      string         `json:"payment_method_type,omitempty"`
	PaymentMethod          string         `json:"payment_method,omitempty"`
	AuthType               string         `json:"auth_type,omitempty"`
	Card                   *Card          `json:"card,omitempty"`
	TxnDetail              *TxnDetail     `json:"txn_detail,omitempty"`
	PaymentGatewayResponse map[string]any `json:"payment_gateway_response,omitempty"`
	EmiDetails             map[string]any `json:"emi_details,omitempty"`
	GatewayID              int            `json:"gateway_id,omitempty"`
	GatewayReferenceID     any            `json:"gateway_reference_id,omitempty"`
	RespCode               any            `json:"resp_code,omitempty"`
	RespMessage            any            `json:"resp_message,omitempty"`
	RespCategory           any            `json:"resp_category,omitempty"`
	BankErrorCode          string         `json:"bank_error_code,omitempty"`
	BankErrorMessage       string         `json:"bank_error_message,omitempty"`
	Refunds                []OrderRefund  `json:"refunds,omitempty"`
	Offers                 []any          `json:"offers,omitempty"`

	// PARTIAL_CHARGED only — list of transactions
	TxnList []OrderStatusTxn `json:"txn_list,omitempty"`

	// shared
	PaymentLinks   *PaymentLinks  `json:"payment_links,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	AdditionalInfo map[string]any `json:"additional_info,omitempty"`

	// udf
	UDF1  string `json:"udf1,omitempty"`
	UDF2  string `json:"udf2,omitempty"`
	UDF3  string `json:"udf3,omitempty"`
	UDF4  string `json:"udf4,omitempty"`
	UDF5  string `json:"udf5,omitempty"`
	UDF6  string `json:"udf6,omitempty"`
	UDF7  string `json:"udf7,omitempty"`
	UDF8  string `json:"udf8,omitempty"`
	UDF9  string `json:"udf9,omitempty"`
	UDF10 string `json:"udf10,omitempty"`
}

// OrderStatusTxn is one entry in OrderStatusResponse.TxnList (PARTIAL_CHARGED).
type OrderStatusTxn struct {
	TxnID                  string         `json:"txn_id,omitempty"`
	TxnUUID                string         `json:"txn_uuid,omitempty"`
	TxnIntentID            string         `json:"txn_intent_id,omitempty"`
	Status                 string         `json:"status,omitempty"`
	PaymentMethod          string         `json:"payment_method,omitempty"`
	PaymentMethodType      string         `json:"payment_method_type,omitempty"`
	AuthType               string         `json:"auth_type,omitempty"`
	Card                   *Card          `json:"card,omitempty"`
	Refunded               bool           `json:"refunded,omitempty"`
	AmountRefunded         float64        `json:"amount_refunded,omitempty"`
	EffectiveAmount        float64        `json:"effective_amount,omitempty"`
	Refunds                []OrderRefund  `json:"refunds,omitempty"`
	RespCode               any            `json:"resp_code,omitempty"`
	RespMessage            any            `json:"resp_message,omitempty"`
	RespCategory           any            `json:"resp_category,omitempty"`
	BankErrorCode          string         `json:"bank_error_code,omitempty"`
	BankErrorMessage       string         `json:"bank_error_message,omitempty"`
	TxnDetail              *TxnDetail     `json:"txn_detail,omitempty"`
	PaymentGatewayResponse map[string]any `json:"payment_gateway_response,omitempty"`
	EmiDetails             map[string]any `json:"emi_details,omitempty"`
	GatewayID              int            `json:"gateway_id,omitempty"`
	GatewayReferenceID     any            `json:"gateway_reference_id,omitempty"`
	Offers                 []any          `json:"offers,omitempty"`
	Metadata               map[string]any `json:"metadata,omitempty"`
}

// TxnDetail is the txn_detail block returned in order status, refund, and webhook responses.
type TxnDetail struct {
	OrderID                   string             `json:"order_id,omitempty"`
	TxnID                     string             `json:"txn_id,omitempty"`
	TxnUUID                   string             `json:"txn_uuid,omitempty"`
	TxnIntentID               string             `json:"txn_intent_id,omitempty"`
	Status                    string             `json:"status,omitempty"`
	TxnAmount                 float64            `json:"txn_amount,omitempty"`
	NetAmount                 float64            `json:"net_amount,omitempty"`
	SurchargeAmount           any                `json:"surcharge_amount,omitempty"`
	TaxAmount                 any                `json:"tax_amount,omitempty"`
	OfferDeductionAmount      any                `json:"offer_deduction_amount,omitempty"`
	RemainingRefundableAmount float64            `json:"remaining_refundable_amount,omitempty"`
	Currency                  string             `json:"currency,omitempty"`
	TxnFlowType               string             `json:"txn_flow_type,omitempty"`
	Gateway                   string             `json:"gateway,omitempty"`
	GatewayID                 int                `json:"gateway_id,omitempty"`
	MerchantIdentifier        string             `json:"merchant_identifier,omitempty"`
	BankName                  string             `json:"bank_name,omitempty"`
	IsCvvLessTxn              bool               `json:"is_cvv_less_txn,omitempty"`
	ExpressCheckout           bool               `json:"express_checkout,omitempty"`
	Redirect                  bool               `json:"redirect,omitempty"`
	Created                   string             `json:"created,omitempty"`
	LastUpdated               string             `json:"last_updated,omitempty"`
	ErrorCode                 any                `json:"error_code,omitempty"`
	ErrorMessage              string             `json:"error_message,omitempty"`
	Metadata                  map[string]any     `json:"metadata,omitempty"`
	TxnAmountBreakup          []TxnAmountBreakup `json:"txn_amount_breakup,omitempty"`
}

type TxnAmountBreakup struct {
	Name   string  `json:"name,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	SNo    int     `json:"sno,omitempty"`
	Method string  `json:"method,omitempty"`
}

// OrderRefund is a single refund entry returned across order status, refund, and webhook responses.
type OrderRefund struct {
	UniqueRequestID          string  `json:"unique_request_id,omitempty"`
	ID                       any     `json:"id,omitempty"`
	Ref                      any     `json:"ref,omitempty"`
	Status                   string  `json:"status,omitempty"`
	Amount                   float64 `json:"amount,omitempty"`
	RefundType               string  `json:"refund_type,omitempty"`
	RefundSource             string  `json:"refund_source,omitempty"`
	InitiatedBy              string  `json:"initiated_by,omitempty"`
	SentToGateway            bool    `json:"sent_to_gateway,omitempty"`
	Created                  string  `json:"created,omitempty"`
	LastUpdated              string  `json:"last_updated,omitempty"`
	ExpectedRefundCreditTime string  `json:"expected_refund_credit_time,omitempty"`
	ErrorCode                any     `json:"error_code,omitempty"`
	ErrorMessage             string  `json:"error_message,omitempty"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Webhook
// ──────────────────────────────────────────────────────────────────────────────

type WebhookEvent struct {
	EventName   string         `json:"event_name"`
	DateCreated string         `json:"date_created"`
	Content     WebhookContent `json:"content"`
	ID          string         `json:"id,omitempty"`
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
	PaymentMethod     string        `json:"payment_method,omitempty"`
	Card              *Card         `json:"card,omitempty"`
	Refunds           []OrderRefund `json:"refunds,omitempty"`
	UDF1              string        `json:"udf1,omitempty"`

	TxnList                  []WebhookTxn   `json:"txn_list,omitempty"`
	EffectiveAmount          float64        `json:"effective_amount,omitempty"`
	PaidAmount               any            `json:"paid_amount,omitempty"`
	CustomerEmail            string         `json:"customer_email,omitempty"`
	CustomerPhoneCountryCode *string        `json:"customer_phone_country_code,omitempty"`
	AdditionalInfo           map[string]any `json:"additional_info,omitempty"`
	LastUpdated              string         `json:"last_updated,omitempty"`
	PaymentLinks             *PaymentLinks  `json:"payment_links,omitempty"`
	Metadata                 map[string]any `json:"metadata,omitempty"`
	DateCreated              string         `json:"date_created,omitempty"`
	CustomerID               string         `json:"customer_id,omitempty"`
}

type WebhookTxn struct {
	PaymentMethod          string         `json:"payment_method,omitempty"`
	PaymentMethodType      string         `json:"payment_method_type,omitempty"`
	PaymentGatewayResponse map[string]any `json:"payment_gateway_response,omitempty"`
	GatewayReferenceID     *string        `json:"gateway_reference_id,omitempty"`
	TxnDetail              map[string]any `json:"txn_detail,omitempty"`
	Status                 string         `json:"status,omitempty"`
	Offers                 []any          `json:"offers,omitempty"`
	RespMessage            *string        `json:"resp_message,omitempty"`
	Refunded               bool           `json:"refunded,omitempty"`
	GatewayID              int            `json:"gateway_id,omitempty"`
	TxnIntentID            string         `json:"txn_intent_id,omitempty"`
	EffectiveAmount        float64        `json:"effective_amount,omitempty"`
	AuthType               string         `json:"auth_type,omitempty"`
	RespCode               *string        `json:"resp_code,omitempty"`
	Card                   *Card          `json:"card,omitempty"`
	TxnUUID                string         `json:"txn_uuid,omitempty"`
	RespCategory           *string        `json:"resp_category,omitempty"`
	BankErrorCode          string         `json:"bank_error_code,omitempty"`
	Metadata               map[string]any `json:"metadata,omitempty"`
	EmiDetails             any            `json:"emi_details,omitempty"`
}
