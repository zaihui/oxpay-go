package oxpay

const (
	TransactionTypeSale          string = "200"
	TransactionTypeVoid          string = "300"
	TransactionTypeRefund        string = "400"
	TransactionTypeEWalletRefund string = "400"
	TransactionTypePreAuth       string = "100"
	TransactionTypeCapture       string = "220"
	TransactionTypeEWalletVoid   string = "81002"
)

const (
	TransactionStatusRequested  string = "1"
	TransactionStatusOK         string = "2"
	TransactionStatusAuthorized string = "3"
	TransactionStatusReversed   string = "4"
	TransactionStatusDenied     string = "5"
	TransactionStatusRefunded   string = "11"
	TransactionStatusVoided     string = "12"
	TransactionStatusCaptured   string = "16"
	TransactionStatusRetrieval  string = "17"
	TransactionStatusUserPaying string = "71"
	TransactionStatusPending    string = "71"
	TransactionStatusClosed     string = "78"
)
const (
	ResponseCodeOK                       string = "00"
	ResponseCodeAuthTerminalParamInvalid string = "1211"

	// Sale
	ResponseCodeSaleTerminalDisabled       string = "2005"
	ResponseCodeSaleParamInvalidOrNotFound string = "2011"
	ResponseCodeSaleInvalidTerminal        string = "2018"
	ResponseCodeSaleDenied                 string = "2090"

	// Void
	ResponseCodeVoidTerminalDisable        string = "2105"
	ResponseCodeVoidParamInvalidOrNotFound string = "2111"
	ResponseCodeVoidInvalidTerminal        string = "2118"
	ResponseCodeVoidDenied                 string = "2190"

	// Refund
	ResponseCodeRefundTerminalDisabled       string = "3205"
	ResponseCodeRefundParamInvalidOrNotFound string = "3211"
	ResponseCodeRefundInvalidTerminal        string = "3218"
	ResponseCodeRefundDenied                 string = "3290"

	// QueryDetail
	ResponseCodeQueryDetailTerminalDisabled       string = "4005"
	ResponseCodeQueryDetailParamInvalidOrNotFound string = "4011"
	ResponseCodeQueryDetailInvalidTerminal        string = "4018"
	ResponseCodeQueryDetailFailNotFoundInGW2      string = "4019"
	ResponseCodeQueryDetailFailByNoDataMatch      string = "4090"

	// EWallet Void
	ResponseCodeEWalletVoidTerminalDisabled       string = "7805"
	ResponseCodeEWalletVoidParamInvalidOrNotFound string = "7811"
	ResponseCodeEWalletVoidInvalidTerminal        string = "7818"
	ResponseCodeEWalletVoidFailed                 string = "7890"

	// All
	ResponseCodeAllJsonObjectMissing string = "9910"
	ResponseCodeAllJsonObjectErr     string = "9911"
	ResponseCodeAllUnknownErr        string = "9999"
)
