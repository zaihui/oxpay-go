// Package oxpay provides the binding for OxPay REST APIs.
package oxpay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

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

// defaultHTTPTimeout is the default timeout on the http.Client used by the library.
// This is chosen to be consistent with the other Stripe language libraries and
// to coordinate with other timeouts configured in the Stripe infrastructure.
const defaultHTTPTimeout = 80 * time.Second

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
type APIResponse struct {
	RawJSON    []byte
	Status     string
	StatusCode int
}

// APIResource is a type assigned to structs that may come from Stripe API
// endpoints and contains facilities common to all of them.
type APIResource struct {
	LastResponse *APIResponse `json:"-"`
}

// SetLastResponse sets the HTTP response that returned the API resource.
func (r *APIResource) SetLastResponse(response *APIResponse) {
	r.LastResponse = response
}
func newAPIResponse(res *http.Response) (*APIResponse, error) {
	response := &APIResponse{
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	response.RawJSON = resBody
	return response, nil
}

type Backend interface {
	Call(method, path, key string, params ParamsContainer, v LastResponseSetter) error
	CallRaw(method, path, key string, params *Params, v LastResponseSetter) error
	//CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *Params, v LastResponseSetter) error
	//SetMaxNetworkRetries(maxNetworkRetries int64)
}

// BackendConfig is used to configure a new Stripe backend.
type BackendConfig struct {
	// EnableTelemetry allows request metrics (request id and duration) to be sent
	// to Stripe in subsequent requests via the `X-Stripe-Client-Telemetry` header.
	//
	// This value is a pointer to allow us to differentiate an unset versus
	// empty value. Use stripe.Bool for an easy way to set this value.
	//
	// Defaults to false.
	EnableTelemetry *bool

	// HTTPClient is an HTTP client instance to use when making API requests.
	//
	// If left unset, it'll be set to a default HTTP client for the package.
	HTTPClient *http.Client
	// MaxNetworkRetries sets maximum number of times that the library will
	// retry requests that appear to have failed due to an intermittent
	// problem.
	//
	// This value is a pointer to allow us to differentiate an unset versus
	// empty value. Use stripe.Int64 for an easy way to set this value.
	//
	// Defaults to DefaultMaxNetworkRetries (2).
	MaxNetworkRetries *int64

	// URL is the base URL to use for API paths.
	//
	// This value is a pointer to allow us to differentiate an unset versus
	// empty value. Use stripe.String for an easy way to set this value.
	//
	// If left empty, it'll be set to the default for the SupportedBackend.
	URL *string
}

// BackendImplementation is the internal implementation for making HTTP calls
// to Stripe.
//
// The public use of this struct is deprecated. It will be unexported in a
// future version.
type BackendImplementation struct {
	Type       SupportedBackend
	URL        string
	HTTPClient *http.Client
}

// Call is the Backend.Call implementation for invoking Stripe APIs.
func (s *BackendImplementation) Call(method, path, mcptid string, params ParamsContainer, v LastResponseSetter) error {
	// todo 这里可能有个坑
	return s.CallRaw(method, path, mcptid, params.GetParams(), v)
}

// CallRaw is the implementation for invoking Stripe APIs internally without a backend.
func (s *BackendImplementation) CallRaw(method, path, mcptid string, params *Params, v LastResponseSetter) error {
	data, err := json.Marshal(params)
	if err != nil {
		return err
	}
	bodyBuffer := bytes.NewReader(data)

	req, err := s.NewRequest(method, path, mcptid, "application/json", bodyBuffer)
	if err != nil {
		return err
	}
	response, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	apiResponse, err := newAPIResponse(response)
	if err != nil {
		return err
	}
	v.SetLastResponse(apiResponse)
	return nil
}
func (s *BackendImplementation) NewRequest(method, path, mcptid, contentType string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = s.URL + path

	req, err := http.NewRequest(method, path, body)
	req.Header.Add("Content-Type", contentType)
	req.Header.Set("Accept", contentType)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//
// Private variables
//

var backends Backends

var httpClient = &http.Client{
	Timeout: defaultHTTPTimeout,
}

// Backends are the currently supported endpoints.
type Backends struct {
	API, Connect, Uploads Backend
	mu                    sync.RWMutex
}

// GetBackend returns one of the library's supported backends based off of the
// given argument.
//
// It returns an existing default backend if one's already been created.
func GetBackend(backendType SupportedBackend) Backend {
	var backend Backend

	backends.mu.RLock()
	switch backendType {
	case APIBackend:
		backend = backends.API
	}
	backends.mu.RUnlock()
	if backend != nil {
		return backend
	}
	backend = GetBackendWithConfig(
		backendType,
		&BackendConfig{
			HTTPClient:        httpClient,
			MaxNetworkRetries: nil, // Set by GetBackendWithConfiguation when nil
			URL:               nil, // Set by GetBackendWithConfiguation when nil
		},
	)

	SetBackend(backendType, backend)
	return backend
}

// GetBackendWithConfig is the same as GetBackend except that it can be given a
// configuration struct that will configure certain aspects of the backend
// that's return.
func GetBackendWithConfig(backendType SupportedBackend, config *BackendConfig) Backend {
	if config.HTTPClient == nil {
		config.HTTPClient = httpClient
	}

	switch backendType {
	case APIBackend:
		if config.URL == nil {
			config.URL = String(APIURL)
		}
		return newBackendImplementation(backendType, config)
	}
	return nil
}

func newBackendImplementation(backendType SupportedBackend, config *BackendConfig) Backend {
	return &BackendImplementation{
		HTTPClient: config.HTTPClient,
		Type:       backendType,
		URL:        *config.URL,
	}
}

// SetBackend sets the backend used in the binding.
func SetBackend(backend SupportedBackend, b Backend) {
	backends.mu.Lock()
	defer backends.mu.Unlock()

	switch backend {
	case APIBackend:
		backends.API = b
	}
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}
