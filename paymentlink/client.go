package paymentlink

import (
	"github.com/li31727/oxpay-go"
	"net/http"
)

// Client is used to invoke /payment_links APIs.
type Client struct {
	B          oxpay.Backend
	McpTID     string
	ApiBackend oxpay.SupportedBackend
}

// New creates a new payment link.
func New(params *oxpay.PaymentLinkParams) (*oxpay.PaymentLink, error) {
	return getC().New(params)
}

func getC() Client {
	return Client{oxpay.GetBackend(oxpay.APIBackend), oxpay.McpTID, oxpay.APIBackend}
}
func (c Client) getPath(relativePath string) string {
	return "/" + string(c.ApiBackend) + relativePath
}

// New creates a new payment link.
func (c Client) New(params *oxpay.PaymentLinkParams) (*oxpay.PaymentLink, error) {
	paymentlink := &oxpay.PaymentLink{}
	err := c.B.Call(
		http.MethodPost,
		c.getPath("/v6/payment"),
		c.McpTID,
		params,
		paymentlink,

	)
	paymentlink.PayLink = string(paymentlink.LastResponse.RawJSON)
	return paymentlink, err
}
