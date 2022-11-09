package refund

import (
	"errors"
	"github.com/li31727/oxpay-go/v5"
	"net/http"
)

// Client is used to invoke /payment_links APIs.
type Client struct {
	B          oxpay.Backend
	McpTID     string
	ApiBackend oxpay.SupportedBackend
}

// New creates a new refund.
func New(id string, params *oxpay.RefundParams, header *oxpay.Head) (*oxpay.Refund, error) {
	return getC().New(id, params, header)
}

// New a refund, card refund not support PAYNOW and APMs,ewallet refund just support by some ewallet,e.g:wechat_pay
func (c Client) New(id string, params *oxpay.RefundParams, header *oxpay.Head) (*oxpay.Refund, error) {
	refund := &oxpay.Refund{}
	if params == nil {
		params = &oxpay.RefundParams{}
	}
	params.TradeNo = id
	p := oxpay.GetParams(params, header)
	path, err := c.chooseUrlByPaymentMethod(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(
		http.MethodPost,
		c.getPath(path),
		c.McpTID,
		p,
		refund,
	)
	return refund, err

}
func (c Client) chooseUrlByPaymentMethod(paymentIntent *oxpay.RefundParams) (string, error) {
	var path string
	switch paymentIntent.PaymentMethod {
	case "ewallet":
		path = "/v5/ewallet/refund"
		if paymentIntent.EType == "" || paymentIntent.Currency == "" {
			return "", errors.New("the eType or currency can not empty")
		}
	case "card":
		path = "/v5/refund"
	}
	return path, nil
}

func (c Client) getPath(relativePath string) string {
	return "/" + string(c.ApiBackend) + relativePath
}

func getC() Client {
	return Client{oxpay.GetBackend(oxpay.APIBackend), oxpay.McpTID, oxpay.APIBackend}
}
