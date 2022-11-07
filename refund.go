package oxpay

// RefundParams PaymentLinkParams Creates a PaymentLinkParams object.
//
// Create a pay link ,then return the pay url for browser.
// todo 退款参数
type RefundParams struct {
	Params `form:"*"`
}
