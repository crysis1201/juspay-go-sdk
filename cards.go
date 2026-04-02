package juspay

import (
	"context"
	"encoding/json"
	"fmt"
)

// ListCards fetches saved cards for a customer.
// GET /cards
func (c *Client) ListCards(ctx context.Context, auth MerchantAuth, customerID string) (*CardsResponse, error) {
	path := fmt.Sprintf("/cards?options.check_cvv_less_support=true&customer_id=%s", customerID)
	respBody, err := c.doRequest(ctx, "GET", path, nil, auth)
	if err != nil {
		return nil, err
	}

	var resp CardsResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FindCardToken looks up a card by its card_reference and returns the card_token.
// Returns an empty string if no matching card is found.
func FindCardToken(cards *CardsResponse, cardReference string) string {
	if cards == nil {
		return ""
	}
	for _, card := range cards.Cards {
		if card.CardReference == cardReference {
			return card.CardToken
		}
	}
	return ""
}
