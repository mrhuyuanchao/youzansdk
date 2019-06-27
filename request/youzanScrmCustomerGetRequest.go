package request

import (
	"fmt"
	"net/url"
)

type YouzanScrmCustomerGetRequest struct {
	AccountID   string `json:"account_id"`
	AccountType string `json:"account_type"`
}

func (y YouzanScrmCustomerGetRequest) GetURL() string {
	return "https://open.youzan.com/api/oauthentry"
}
func (y YouzanScrmCustomerGetRequest) GetMethod() string {
	return "GET"
}
func (y YouzanScrmCustomerGetRequest) GetApiName() string {
	return "youzan.scrm.customer.get"
}
func (y YouzanScrmCustomerGetRequest) GetApiVersion() string {
	return "3.1.0"
}
func (y YouzanScrmCustomerGetRequest) GetParam() map[string]string {
	param := make(map[string]string, 0)
	content := fmt.Sprintf("{\"account_type\":\"%s\", \"account_id\":\"%s\"}", y.AccountType, y.AccountID)
	param["account"] = content
	return param
}
func (y YouzanScrmCustomerGetRequest) GetBodyParam() interface{} {
	return nil
}

func (y YouzanScrmCustomerGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y YouzanScrmCustomerGetRequest) CheckParam() error {
	return nil
}
