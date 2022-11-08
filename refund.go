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
	ProcessingCode      string `json:"processingCode"`
	GpsLatitude         string `json:"gpsLatitude"`
	GpsLongitude        string `json:"gpsLongitude"`
	ServiceAmount       string `json:"serviceAmount"`
	ServiceRate         string `json:"serviceRate"`
	GstAmount           string `json:"gstAmount"`
	GstRate             string `json:"gstRate"`
}
