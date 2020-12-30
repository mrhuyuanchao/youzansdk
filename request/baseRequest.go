package request

import "net/url"

const youzanApiBaseUrl = "https://open.youzanyun.com"
const apiRequestPath = "%s/api/%s/%s?access_token=%s"
const apiRequestPathAuthExempt = "%s/api/auth_exempt/%s/%s"
const apiRequestPathTextarea = "%s/api/_textarea_/%s/%s?access_token=%s"
const apiRequestPathToken = "%s/auth/token"

// BaseRequest ..
type BaseRequest interface {
	GetMethod() string
	GetApiName() string
	GetApiVersion() string
	GetParam() map[string]string
	GetBodyParam() interface{}
	CheckParam() error
	GetUrlValues() url.Values
	GetRequestUrl(token string) string
}
