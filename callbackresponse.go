package oxpay

type Data struct {
	AmountInCNY             string `json:"amountInCNY"`
	AmountInCurrencySettled string `json:"amountInCurrencySettled"`
	CurrencySettled         string `json:"currencySettled"`
	EWalletType             string `json:"eWalletType"`
	DiscountAmount          string `json:"discountAmount"`
	PaidAmount              string `json:"paidAmount"`
}

type CallbackResponse struct {
	McpTID           string `json:"mcptid"`
	Currency         string `json:"currency"`
	TotalAmount      string `json:"totalAmount"`
	ReferenceNo      string `json:"referenceNo"`
	TransactionId    string `json:"transactionId"`
	TransactionState string `json:"transactionState"`
	Stan             string `json:"stan"`
	ReceiptNumber    string `json:"receiptNumber"`
	CardHolderName   string `json:"cardHolderName"`
	TruncatedPan     string `json:"truncatedPan"`
	BrandName        string `json:"brandName"`
	ResponseCode     string `json:"responseCode"`
	ResponseMsg      string `json:"responseMsg"`
	Data             Data   `json:"data"`
}
