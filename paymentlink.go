package oxpay

// PaymentLinkParams Creates a PaymentLinkParams object.
//
// Create a pay link ,then return the pay url for browser.
// todo 创建支付链接参数
type PaymentLinkParams struct {
	*Params
	MCPTID        string `json:"mcptid"`
	Currency      string `json:"currency"`
	TotalAmount   string `json:"totalAmount"` //单位分
	ReferenceNo   string `json:"referenceNo"`
	StatusUrl     string `json:"statusUrl"`
	ReturnUrl     string `json:"returnUrl"`
	ItemDetail    string `json:"itemDetail"`
	CustomerEmail string `json:"customerEmail"`
	Tokenize      string `json:"tokenize"`
}

type PaymentLink struct {
	APIResource
}
