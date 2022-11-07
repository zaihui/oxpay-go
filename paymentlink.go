package oxpay

// PaymentLinkParams Creates a PaymentLinkParams object.
//
// Create a pay link ,then return the pay url for browser.
// todo 创建支付链接参数
type PaymentLinkParams struct {
	Params `form:"*"`
}
