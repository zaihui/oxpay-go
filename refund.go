package oxpay

// RefundParams PaymentLinkParams Creates a PaymentLinkParams object.
//
// Create a pay link ,then return the pay url for browser.
// todo 退款参数
type RefundParams struct {
	Params                `form:"*"`
	PaymentMethod         string `json:"payment_method"` //card or ewallet
	Currency              string `json:"currency"`
	TradeNo               string `json:"tradeNo"`
	EType                 string `json:"eType"`
	TotalAmount           string `json:"totalAmount"`
	OriginalTransactionId string `json:"originalTransactionId"` //  paymentintent response`s origTnxId
}

type Refund struct {
	APIResource
}
