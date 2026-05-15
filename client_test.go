package juspay

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServer(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	ts := httptest.NewServer(handler)
	client := NewClient(Config{BaseURL: ts.URL})
	return client, ts
}

func TestCreateSession(t *testing.T) {
	client, ts := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/session" {
			t.Errorf("expected /session, got %s", r.URL.Path)
		}

		// Verify auth
		user, _, ok := r.BasicAuth()
		if !ok || user != "test-api-key" {
			t.Errorf("expected basic auth with user test-api-key, got %s", user)
		}
		if r.Header.Get("x-merchantid") != "TEST_MERCHANT" {
			t.Errorf("expected x-merchantid TEST_MERCHANT, got %s", r.Header.Get("x-merchantid"))
		}

		// Verify request body
		var req CreateSessionRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.OrderID != "ORDER-123" {
			t.Errorf("expected order_id ORDER-123, got %s", req.OrderID)
		}
		if req.Amount != "100.00" {
			t.Errorf("expected amount 100.00, got %s", req.Amount)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(SessionResponse{
			OrderID: "ORDER-123",
			Status:  "NEW",
			PaymentLinks: &PaymentLinks{
				Web: strPtr("https://example.com/pay"),
			},
		})
	})
	defer ts.Close()

	auth := MerchantAuth{APIKey: "test-api-key", MerchantID: "TEST_MERCHANT"}
	resp, err := client.CreateSession(context.Background(), auth, CreateSessionRequest{
		OrderID:  "ORDER-123",
		Amount:   "100.00",
		Currency: "INR",
		Action:   "paymentPage",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.OrderID != "ORDER-123" {
		t.Errorf("expected ORDER-123, got %s", resp.OrderID)
	}
	if resp.Status != "NEW" {
		t.Errorf("expected NEW, got %s", resp.Status)
	}
	if resp.PaymentLinks == nil || resp.PaymentLinks.Web == nil {
		t.Fatal("expected payment links with web URL")
	}
	if *resp.PaymentLinks.Web != "https://example.com/pay" {
		t.Errorf("expected https://example.com/pay, got %s", *resp.PaymentLinks.Web)
	}
}

func TestListCards(t *testing.T) {
	client, ts := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Query().Get("customer_id") != "cust-1" {
			t.Errorf("expected customer_id=cust-1, got %s", r.URL.Query().Get("customer_id"))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CardsResponse{
			Cards: []Card{
				{CardToken: "tok_123", CardReference: "ref_abc", LastFourDigits: "4242"},
				{CardToken: "tok_456", CardReference: "ref_def", LastFourDigits: "1234"},
			},
		})
	})
	defer ts.Close()

	auth := MerchantAuth{APIKey: "test-key", MerchantID: "MERCHANT"}
	resp, err := client.ListCards(context.Background(), auth, "cust-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Cards) != 2 {
		t.Fatalf("expected 2 cards, got %d", len(resp.Cards))
	}
	if resp.Cards[0].CardToken != "tok_123" {
		t.Errorf("expected tok_123, got %s", resp.Cards[0].CardToken)
	}
}

func TestFindCardToken(t *testing.T) {
	cards := &CardsResponse{
		Cards: []Card{
			{CardToken: "tok_123", CardReference: "ref_abc"},
			{CardToken: "tok_456", CardReference: "ref_def"},
		},
	}

	if tok := FindCardToken(cards, "ref_abc"); tok != "tok_123" {
		t.Errorf("expected tok_123, got %s", tok)
	}
	if tok := FindCardToken(cards, "ref_def"); tok != "tok_456" {
		t.Errorf("expected tok_456, got %s", tok)
	}
	if tok := FindCardToken(cards, "ref_unknown"); tok != "" {
		t.Errorf("expected empty, got %s", tok)
	}
	if tok := FindCardToken(nil, "ref_abc"); tok != "" {
		t.Errorf("expected empty for nil, got %s", tok)
	}
}

func TestCreateTransaction(t *testing.T) {
	client, ts := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/txns" {
			t.Errorf("expected POST /txns, got %s %s", r.Method, r.URL.Path)
		}

		var req CreateTransactionRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.CardToken != "tok_123" {
			t.Errorf("expected card_token tok_123, got %s", req.CardToken)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(TransactionResponse{
			TxnID:  "txn_789",
			Status: "CHARGED",
		})
	})
	defer ts.Close()

	auth := MerchantAuth{APIKey: "key", MerchantID: "M"}
	resp, raw, err := client.CreateTransaction(context.Background(), auth, CreateTransactionRequest{
		CardToken: "tok_123",
		OrderID:   "ORD-1",
		Amount:    "50.00",
		Currency:  "INR",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.TxnID != "txn_789" {
		t.Errorf("expected txn_789, got %s", resp.TxnID)
	}
	if resp.Status != "CHARGED" {
		t.Errorf("expected CHARGED, got %s", resp.Status)
	}
	if len(raw) == 0 {
		t.Error("expected non-empty raw response")
	}
}

func TestCreateRefund(t *testing.T) {
	client, ts := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/orders/ORD-1/refunds" {
			t.Errorf("expected POST /orders/ORD-1/refunds, got %s %s", r.Method, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(RefundResponse{
			ID:     "ref_001",
			Status: "PENDING",
			Amount: 50.00,
		})
	})
	defer ts.Close()

	auth := MerchantAuth{APIKey: "key", MerchantID: "M"}
	resp, err := client.CreateRefund(context.Background(), auth, "ORD-1", CreateRefundRequest{
		OrderID:         "ORD-1",
		Amount:          "50.00",
		UniqueRequestID: "uniq-1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ID != "ref_001" {
		t.Errorf("expected ref_001, got %s", resp.ID)
	}
}

func TestCreateOrder(t *testing.T) {
	client, ts := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/orders" {
			t.Errorf("expected POST /orders, got %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(OrderResponse{OrderID: "ORD-NEW", Status: "CREATED"})
	})
	defer ts.Close()

	auth := MerchantAuth{APIKey: "key", MerchantID: "M"}
	resp, err := client.CreateOrder(context.Background(), auth, CreateOrderRequest{
		OrderID: "ORD-NEW", Amount: "200.00", Currency: "INR",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.OrderID != "ORD-NEW" {
		t.Errorf("expected ORD-NEW, got %s", resp.OrderID)
	}
}

func TestGetOrderStatus(t *testing.T) {
	client, ts := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/orders/ORD-1" {
			t.Errorf("expected GET /orders/ORD-1, got %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(OrderStatusResponse{
			OrderID:  "ORD-1",
			Status:   "CHARGED",
			StatusID: 21,
			Amount:   100,
			Currency: "INR",
		})
	})
	defer ts.Close()

	auth := MerchantAuth{APIKey: "key", MerchantID: "M"}
	resp, err := client.GetOrderStatus(context.Background(), auth, "ORD-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "CHARGED" {
		t.Errorf("expected CHARGED, got %s", resp.Status)
	}
	if resp.StatusID != 21 {
		t.Errorf("expected 21, got %d", resp.StatusID)
	}
}

func TestAPIError(t *testing.T) {
	client, ts := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"invalid_api_key"}`))
	})
	defer ts.Close()

	auth := MerchantAuth{APIKey: "bad-key", MerchantID: "M"}
	_, err := client.CreateSession(context.Background(), auth, CreateSessionRequest{})
	if err == nil {
		t.Fatal("expected error for 401 response")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 401 {
		t.Errorf("expected status 401, got %d", apiErr.StatusCode)
	}
}

func TestParseWebhookEvent(t *testing.T) {
	payload := `{
		"event_name": "ORDER_SUCCEEDED",
		"date_created": "2024-01-01T00:00:00Z",
		"content": {
			"order": {
				"order_id": "ORD-123",
				"merchant_id": "TESTZUZU",
				"status": "CHARGED",
				"status_id": 21,
				"amount": 100.00,
				"currency": "INR",
				"txn_id": "txn_abc",
				"payment_method_type": "CARD"
			}
		}
	}`

	event, err := ParseWebhookEvent([]byte(payload))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if event.EventName != EventOrderSucceeded {
		t.Errorf("expected ORDER_SUCCEEDED, got %s", event.EventName)
	}
	if event.Content.Order.OrderID != "ORD-123" {
		t.Errorf("expected ORD-123, got %s", event.Content.Order.OrderID)
	}
	if event.Content.Order.StatusID != 21 {
		t.Errorf("expected 21, got %d", event.Content.Order.StatusID)
	}
}

func TestParseWebhookEventInvalid(t *testing.T) {
	_, err := ParseWebhookEvent([]byte("not json"))
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func strPtr(s string) *string {
	return &s
}
