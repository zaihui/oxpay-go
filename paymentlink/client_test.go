package paymentlink_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/li31727/oxpay-go/v5"
	"github.com/li31727/oxpay-go/v5/paymentlink"
)

const (
	// this is provided by OxPay
	MCPTID = "YOUR MCPTID"

	// this is callback url for OxPay notify transaction result.
	StatusUrl = "https://xxdfafd"

	// this is front url
	ReturnUrl = "https://xxdfafd"
)

func TestGetPayLink(t *testing.T) {
	oxpay.McpTID = MCPTID
	params := &oxpay.PaymentLinkParams{
		MCPTID:        MCPTID,
		Currency:      "SGD",
		TotalAmount:   "100",
		ReferenceNo:   "test_ReferenceNo" + strconv.Itoa(int(time.Now().Unix())),
		StatusUrl:     StatusUrl,
		ReturnUrl:     ReturnUrl,
		ItemDetail:    "test_item",
		CustomerEmail: "xxx@qq.com",
		Tokenize:      "Y",
	}
	link, err := paymentlink.New(params, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(link.PayLink)
}
