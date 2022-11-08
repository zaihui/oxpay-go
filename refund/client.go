package refund

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

// New creates a new refund.
func New(params *oxpay.RefundParams) (*oxpay.Refund, error) {
	return getC().New(params)
}

// New a refund, card refund not support PAYNOW and APMs,ewallet refund just support by some ewallet,e.g:wechat_pay
func (c Client) New(params *oxpay.RefundParams) (*oxpay.Refund, error) {
	refund := &oxpay.Refund{}
	params.Params.Data = params
	path := ""
	switch params.PaymentMethod {
	case "ewallet":
		path = "/v5/ewallet/refund"
	case "card":
		path = "/v5/refund"
	}
	params.Data = params
	err := c.B.Call(
		http.MethodPost,
		c.getPath(path),
		c.McpTID,
		&params.Params,
		refund,
	)
	return refund, err

}

func (c Client) getPath(relativePath string) string {
	return "/" + string(c.ApiBackend) + relativePath
}

func getC() Client {
	return Client{oxpay.GetBackend(oxpay.APIBackend), oxpay.McpTID, oxpay.APIBackend}
}
