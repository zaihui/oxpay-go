// Package oxpay provides the binding for OxPay REST APIs.
package oxpay

const (
	// APIVersion is the currently supported API version
	APIVersion string = apiVersion

	// APIBackend is a constant representing the API service backend.
	APIBackend SupportedBackend = "api"

	// APIURL is the URL of the API service backend.
	APIURL string = "https://gw2.mcpayment.net"

	// TestAPIURL is the URL of the test API service backend.
	TestAPIURL = "https://gw2.sandbox.mcpayment.net"
)

// McpTID is the OxPay merchants key used globally in the binding.
var McpTID string

// SupportedBackend is an enumeration of supported OxPay endpoints.
// Currently supported value is "api"
type SupportedBackend string

// LastResponseSetter defines a type that contains an HTTP response from a Stripe
// API endpoint.
type LastResponseSetter interface {
	SetLastResponse(response *APIResponse)
}

// APIResponse encapsulates some common features of a response from the
// Stripe API.
// todo  先占个位，等功能接完后再抽象出成员
type APIResponse struct{}

type Backend interface {
	Call(method, path, key string, params ParamsContainer, v LastResponseSetter) error
	CallRaw(method, path, key string, params *Params, v LastResponseSetter) error
	//CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *Params, v LastResponseSetter) error
	//SetMaxNetworkRetries(maxNetworkRetries int64)
}

//
// Private variables
//

var backends Backends

// Backends are the currently supported endpoints.
type Backends struct {
	API, Connect, Uploads Backend
}

// GetBackend returns one of the library's supported backends based off of the
// given argument.
//
// It returns an existing default backend if one's already been created.
func GetBackend(backendType SupportedBackend) Backend {
	var backend Backend

	switch backendType {
	case APIBackend:
		backend = backends.API
	}
	if backend != nil {
		return backend
	}
	return backend
}
