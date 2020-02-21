package request

import "net/url"

// BaseRequest ..
type BaseRequest interface {
	GetMethod() string
	GetApiName() string
	GetApiVersion() string
	GetParam() map[string]string
	GetBodyParam() interface{}
	CheckParam() error
	GetUrlValues() url.Values
}
