# juspay-go

Go client library for the [Juspay](https://juspay.in) payments API.

## Installation

```bash
go get github.com/crysis1201/juspay-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    juspay "github.com/crysis1201/juspay-go"
)

func main() {
    client := juspay.NewClient(juspay.Config{
        BaseURL: "https://api.juspay.in",
    })

    auth := juspay.MerchantAuth{
        APIKey:     "your-api-key",
        MerchantID: "your-merchant-id",
    }

    // Create a payment session
    email := "customer@example.com"
    phone := "+919999999999"
    returnURL := "https://yoursite.com/callback"

    resp, err := client.CreateSession(context.Background(), auth, juspay.CreateSessionRequest{
        OrderID:             "ORDER-001",
        Amount:              "1000.00",
        CustomerID:          "cust-123",
        CustomerEmail:       &email,
        CustomerPhone:       &phone,
        PaymentPageClientID: "your-client-id",
        Action:              "paymentPage",
        ReturnURL:           &returnURL,
        Currency:            "INR",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Order:", resp.OrderID, "Status:", resp.Status)
}
```

## API Coverage

| Method | Endpoint | Description |
|--------|----------|-------------|
| `CreateSession` | `POST /session` | Create a HyperCheckout payment session |
| `CreateOrder` | `POST /orders` | Create a new order |
| `GetOrderStatus` | `GET /orders/{id}` | Fetch order status |
| `CreateRefund` | `POST /orders/{id}/refunds` | Initiate a refund |
| `CreateTransaction` | `POST /txns` | Charge a saved card (MOTO) |
| `ListCards` | `GET /cards` | List saved cards for a customer |
| `FindCardToken` | - | Helper to look up a card token by card reference |
| `ParseWebhookEvent` | - | Parse a Juspay webhook payload |

## Webhook Handling

```go
event, err := juspay.ParseWebhookEvent(payload)
if err != nil {
    log.Fatal(err)
}

switch event.EventName {
case juspay.EventOrderSucceeded:
    // handle successful payment
case juspay.EventOrderFailed:
    // handle failed payment
case juspay.EventOrderRefunded:
    // handle refund
case juspay.EventOrderRefundFailed:
    // handle refund failure
}
```

## Error Handling

Non-2xx responses are returned as `*juspay.APIError`:

```go
resp, err := client.CreateSession(ctx, auth, req)
if err != nil {
    var apiErr *juspay.APIError
    if errors.As(err, &apiErr) {
        fmt.Println("Status:", apiErr.StatusCode)
        fmt.Println("Body:", apiErr.Body)
    }
}
```

## Testing

```bash
go test -v ./...
```

## License

MIT
