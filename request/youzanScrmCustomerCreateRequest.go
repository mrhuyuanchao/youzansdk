package request

import (
	"encoding/json"
	"errors"
	"net/url"
)

type CustomerCreateModel struct {
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
	Name     string `json:"name"`
}

// YouzanScrmCustomerCreateRequest ..
type YouzanScrmCustomerCreateRequest struct {
	Mobile      string
	CreateModel CustomerCreateModel
}

func (y YouzanScrmCustomerCreateRequest) GetMethod() string {
	return "POSTFORM"
}
func (y YouzanScrmCustomerCreateRequest) GetApiName() string {
	return "youzan.scrm.customer.create"
}
func (y YouzanScrmCustomerCreateRequest) GetApiVersion() string {
	return "3.0.0"
}
func (y YouzanScrmCustomerCreateRequest) GetParam() map[string]string {
	param := make(map[string]string, 0)
	return param
}
func (y YouzanScrmCustomerCreateRequest) GetBodyParam() interface{} {
	return nil
}

func (y YouzanScrmCustomerCreateRequest) GetUrlValues() url.Values {
	u := url.Values{}
	u.Set("mobile", y.Mobile)
	content := "{}"
	byteStr, err := json.Marshal(y.CreateModel)
	if err != nil {
		content = string(byteStr)
	}

	u.Set("customer_create", content)
	return u
}

func (y YouzanScrmCustomerCreateRequest) CheckParam() error {
	if y.Mobile == "" {
		return errors.New("手机号不能为空")
	}
	return nil
}
