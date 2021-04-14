package request

import (
	"fmt"
	"net/url"
)

type YouzanUmpCouponsUnfinishedRequest struct {
	Fields string
}

func (y *YouzanUmpCouponsUnfinishedRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanUmpCouponsUnfinishedRequest) GetApiName() string {
	return "youzan.ump.coupons.unfinished.search"
}
func (y *YouzanUmpCouponsUnfinishedRequest) GetApiVersion() string {
	return "3.0.0"
}
func (y *YouzanUmpCouponsUnfinishedRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanUmpCouponsUnfinishedRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.Fields != "" {
		param["fields"] = y.Fields
	}
	return param
}

func (y *YouzanUmpCouponsUnfinishedRequest) CheckParam() error {
	return nil
}
func (y *YouzanUmpCouponsUnfinishedRequest) GetUrlValues() url.Values {
	return nil
}

func (y *YouzanUmpCouponsUnfinishedRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
