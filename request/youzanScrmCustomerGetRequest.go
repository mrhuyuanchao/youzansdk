package request

import (
	"fmt"
	"net/url"
)

type YouzanScrmCustomerGetRequest struct {
	AccountID   string `json:"account_id"`
	AccountType string `json:"account_type"`
}

func (y YouzanScrmCustomerGetRequest) GetMethod() string {
	return "POST"
}
func (y YouzanScrmCustomerGetRequest) GetApiName() string {
	return "youzan.scrm.customer.get"
}
func (y YouzanScrmCustomerGetRequest) GetApiVersion() string {
	return "3.1.0"
}
func (y YouzanScrmCustomerGetRequest) GetParam() map[string]string {
	return nil
}
func (y YouzanScrmCustomerGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	content := fmt.Sprintf("{\"account_type\":\"%s\", \"account_id\":\"%s\"}", y.AccountType, y.AccountID)
	param["account"] = content
	return param
}

func (y YouzanScrmCustomerGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y YouzanScrmCustomerGetRequest) CheckParam() error {
	return nil
}

func (y YouzanScrmCustomerGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
