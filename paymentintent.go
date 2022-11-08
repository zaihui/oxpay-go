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
	Header              Head   `json:"header"`
	TransactionId       string `json:"transactionId"`
	TransactionType     string `json:"transactionType"`
	TransactionState    string `json:"transactionState"`
	HostResponseDate    string `json:"hostResponseDate"`
	HostResponseMessage string `json:"hostResponseMessage"`
	OrigTnxId           string `json:"origTnxId"`
	ClientRequestDate   string `json:"clientRequestDate"`
	GatewayRequestDate  string `json:"gatewayRequestDate"`
	GatewayResponseDate string `json:"gatewayResponseDate"`
	RespCode            string `json:"respCode"`
	Currency            string `json:"currency"`
	Stan                string `json:"stan"`
	ReceiptNumber       string `json:"receiptNumber"`
	TruncatedPan        string `json:"truncatedPan"`
	BrandName           string `json:"brandName"`
	TotalAmount         string `json:"totalAmount"`
	SalesAmount         string `json:"salesAmount"`
	Rrn                 string `json:"rrn"`
	AuthCode            string `json:"authCode"`
	CardHolderName      string `json:"cardHolderName"`
	CurrencyCode        string `json:"currencyCode"`
	GpsLatitude         string `json:"gpsLatitude"`
	GpsLongitude        string `json:"gpsLongitude"`
	ServiceAmount       string `json:"serviceAmount"`
	ServiceRate         string `json:"serviceRate"`
	GstAmount           string `json:"gstAmount"`
	GstRate             string `json:"gstRate"`
	Email               string `json:"email"`
	Buyer               string `json:"buyer"`
}

// Creates a PaymentIntentCancelParams object.
//
// todo 取消订单参数

type PaymentIntentCancelParams struct {
	Params             `form:"*"`
	CancellationReason *string `form:"cancellation_reason"`
}
