package paymentlink

import (
	"github.com/li31727/oxpay-go/v5"
	"net/http"
)

// Client is used to invoke /payment_links APIs.
type Client struct {
	B          oxpay.Backend
	McpTID     string
	ApiBackend oxpay.SupportedBackend
}

// New creates a new payment link.
func New(params *oxpay.PaymentLinkParams, header *oxpay.Head) (*oxpay.PaymentLink, error) {
	return getC().New(params, header)
}

func getC() Client {
	return Client{oxpay.GetBackend(oxpay.APIBackend), oxpay.McpTID, oxpay.APIBackend}
}
func (c Client) getPath(relativePath string) string {
	return "/" + string(c.ApiBackend) + relativePath
}

// New creates a new payment link.
func (c Client) New(params *oxpay.PaymentLinkParams, header *oxpay.Head) (*oxpay.PaymentLink, error) {
	paymentlink := &oxpay.PaymentLink{}
	p := oxpay.GetParams(params, header)

	err := c.B.Call(
		http.MethodPost,
		c.getPath("/v6/payment"),
		c.McpTID,
		p,
		paymentlink,
	)
	paymentlink.PayLink = string(paymentlink.LastResponse.RawJSON)
	if err != nil && paymentlink.LastResponse.StatusCode != http.StatusOK {
		return nil, err
	}
	return paymentlink, nil
}
