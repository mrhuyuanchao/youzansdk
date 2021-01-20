package request

import (
	"errors"
	"fmt"
	"net/url"
)

// YouzanCrmCustomerPointsGetRequest 获取用户积分
type YouzanCrmCustomerPointsGetRequest struct {
	User struct {
		AccountType int8   `json:"account_type"`
		AccountID   string `json:"account_id"`
	}
}

func (y *YouzanCrmCustomerPointsGetRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanCrmCustomerPointsGetRequest) GetApiName() string {
	return "youzan.crm.customer.points.increase"
}
func (y *YouzanCrmCustomerPointsGetRequest) GetApiVersion() string {
	return "4.0.0"
}
func (y *YouzanCrmCustomerPointsGetRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanCrmCustomerPointsGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	param["user"] = y.User
	return param
}

func (y *YouzanCrmCustomerPointsGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanCrmCustomerPointsGetRequest) CheckParam() error {
	if y.User.AccountType == 0 && y.User.AccountID == "" {
		return errors.New("请指定账号类型和账号ID")
	}
	return nil
}

func (y *YouzanCrmCustomerPointsGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
