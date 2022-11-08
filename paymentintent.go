package oxpay

// Creates a PaymentIntent object.
//

type PaymentIntentParams struct {
	Params               `form:"*"`
	PayMethod            string `json:"pay_method"`
	ID                   string `json:"id"`
	CardToken            string `json:"cardtoken"` //card quick pay token
	Currency             string `json:"currency"`
	TransactionId        string `json:"transactionId"` //OxPay transaction id
	ReferenceNo          string `json:"referenceNo"`
	TotalAmount          string `json:"totalAmount"`
	CustomerEmailAddress string `json:"customerEmailAddress"`
	CustomerName         string `json:"customerName"`
	EType                string `json:"eType"` // pay channel e.g:wechat=77771
}

type PaymentIntent struct {
	APIResource
}

// Creates a PaymentIntentCancelParams object.
//
// todo 取消订单参数

type PaymentIntentCancelParams struct {
	Params             `form:"*"`
	CancellationReason *string `form:"cancellation_reason"`
}
