package paymentlink

import (
	"github.com/li31727/oxpay-go"
)

// Client is used to invoke /payment_links APIs.
type Client struct {
	B   oxpay.Backend
	Key string
}

// New creates a new payment link.
func New(params *oxpay.PaymentLinkParams) (*oxpay.PaymentLink, error) {
	return getC().New(params)
}

func getC() Client {
	return Client{oxpay.GetBackend(oxpay.APIBackend), oxpay.McpTID}
}
