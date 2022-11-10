package refund_test

import (
	"fmt"
	"testing"

	"github.com/li31727/oxpay-go/v5"
	"github.com/li31727/oxpay-go/v5/refund"
)

const (
	// this is provided by OxPay
	MCPTID = "YOUR MCPTID"

	// this is callback url for OxPay notify transaction result.
	StatusUrl = "https://xxdfafd"

	// this is front url
	ReturnUrl = "https://xxdfafd"
)

func TestRefundTransaction(t *testing.T) {
	params := &oxpay.RefundParams{
		PaymentMethod: "ewallet",
		EType:         oxpay.PaymentMethodTypeWECHAT,
		Currency:      "SGD",
		TotalAmount:   "1",
	}
	header := &oxpay.Head{
		AppType:       "W",
		AppVersion:    "spoonxtest.0002.00012.1",
		McpTerminalId: MCPTID,
		Version:       "5",
	}
	id := "7753278"
	o, err := refund.New(id, params, header)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(o.APIResource.LastResponse.RawJSON))
}
