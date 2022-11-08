package oxpay

import (
	"net/url"
)

//
// Public constants
//

type Params struct {
	Header Head `json:"header"`
	Data   any  `json:"data"`
}

// Head is the request common comment.
type Head struct {
	Version       string `json:"version"`       //OxPayAPI version
	AppType       string `json:"appType"`       //app类型 I: iPhone; A: Android; W: Web
	AppVersion    string `json:"appVersion"`    //For example: "AppName.01.20.0" or "WebName.0002.00012.1"
	Status        Status `json:"status"`        //json字符串 只存在于response 的head中
	McpTerminalId string `json:"mcpTerminalId"` // 登录后相应体中获取Mcp Terminal Id mandatory for sale, void, refund, reverse...
	Signature     string `json:"signature"`     //签名 登录后获取 ，并且在每次发送请求时都需要带上
	Uuid          string `json:"uuid"`          // 在pos机上需要设置

}
type Status struct {
	ResponseCode string `json:"responseCode"`
	Message      string `json:"message"`
}

// ParamsContainer is a general interface for which all parameter structs
// should comply. They achieve this by embedding a Params struct and inheriting
// its implementation of this interface.
type ParamsContainer interface {
	GetParams() *Params
}

// GetParams returns a Params struct (itself). It exists because any structs
// that embed Params will inherit it, and thus implement the ParamsContainer
// interface.
func (p *Params) GetParams() *Params {
	return p
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
