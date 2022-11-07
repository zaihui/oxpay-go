package oxpay

// Creates a PaymentIntent object.
//
// todo 下单参数

type PaymentIntentParams struct {
	Params `form:"*"`
}

// Creates a PaymentIntentCancelParams object.
//
// todo 取消订单参数

type PaymentIntentCancelParams struct {
	Params             `form:"*"`
	CancellationReason *string `form:"cancellation_reason"`
}

// Creates a PaymentIntentQueryParams object.
//
// todo 查询订单参数

type PaymentIntentQueryParams struct {
	PaymentIntentParams `form:"*"`
}
