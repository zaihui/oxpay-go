package oxpay

import (
	"context"
	"net/http"
	"net/url"
)

//
// Public constants
//

type Params struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `form:"-"`

	Expand []*string    `form:"expand"`
	Extra  *ExtraValues `form:"*"`

	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header `form:"-"`

	IdempotencyKey *string           `form:"-"` // Passed as header
	Metadata       map[string]string `form:"metadata"`

	// StripeAccount may contain the ID of a connected account. By including
	// this field, the request is made as if it originated from the connected
	// account instead of under the account of the owner of the configured
	// Stripe key.
	StripeAccount *string `form:"-"` // Passed as header
}

// ParamsContainer is a general interface for which all parameter structs
// should comply. They achieve this by embedding a Params struct and inheriting
// its implementation of this interface.
type ParamsContainer interface {
	GetParams() *Params
}

//
// Public types
//

// ExtraValues are extra parameters that are attached to an API request.
// They're implemented as a custom type so that they can have their own
// AppendTo implementation.
type ExtraValues struct {
	url.Values `form:"-"` // See custom AppendTo implementation
}
