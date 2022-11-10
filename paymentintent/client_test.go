package paymentintent_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/li31727/oxpay-go/v5"
	"github.com/li31727/oxpay-go/v5/paymentintent"
)

const (
	// this is provided by OxPay
	MCPTID = ""

	// this is callback url for OxPay notify transaction result.
	StatusUrl = "https://xxdfafd"

	// this is front url
	ReturnUrl = "https://xxdfafd"
)

func TestGetTransactionInfo(t *testing.T) {
	oxpay.McpTID = MCPTID
	header := &oxpay.Head{
		AppType:       "W",
		AppVersion:    "spoonxtest.0002.00012.1",
		McpTerminalId: MCPTID,
		Version:       "5",
	}
	id := "7753278"
	paymentIntent, err := paymentintent.Get(id, nil, header)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(paymentIntent.APIResource.LastResponse.RawJSON))
}

func TestCardQuickPay(t *testing.T) {
	header := &oxpay.Head{
		AppType:       "W",
		AppVersion:    "spoonxtest.0002.00012.1",
		McpTerminalId: MCPTID,
		Version:       "5",
	}
	params := &oxpay.PaymentIntentParams{
		CardToken:     "49c42dd7-9b58-4819-9c97-35fad3d4bb76",
		Currency:      "SGD",
		TotalAmount:   "1",
		ReferenceNo:   "test_ReferenceNo" + strconv.Itoa(int(time.Now().Unix())),
		CustomerEmail: "317727355@qq.com",
		CustomerName:  "li",
	}
	intent, err := paymentintent.New(params, header)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(intent.APIResource.LastResponse.RawJSON))

}
