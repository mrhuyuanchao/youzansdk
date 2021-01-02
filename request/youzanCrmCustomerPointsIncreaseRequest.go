package request

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// YouzanCrmCustomerPointsIncreaseRequest 增加用户积分
type YouzanCrmCustomerPointsIncreaseRequest struct {
	Reason    string
	KdtID     int64
	RequestID string
	ClientID  string
	Points    int64
	User      struct {
		AccountType int8   `json:"account_type"`
		AccountID   string `json:"account_id"`
	}
}

func (y *YouzanCrmCustomerPointsIncreaseRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanCrmCustomerPointsIncreaseRequest) GetApiName() string {
	return "youzan.crm.customer.points.increase"
}
func (y *YouzanCrmCustomerPointsIncreaseRequest) GetApiVersion() string {
	return "4.0.0"
}
func (y *YouzanCrmCustomerPointsIncreaseRequest) GetParam() map[string]string {
	return nil
}
func (y *YouzanCrmCustomerPointsIncreaseRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.RequestID != "" {
		param["biz_value"] = y.RequestID
	} else {
		param["biz_value"] = time.Now().UnixNano()
	}
	param["kdt_id"] = y.KdtID
	param["points"] = y.Points
	param["user"] = y.User
	param["client_id"] = y.ClientID
	return param
}

func (y *YouzanCrmCustomerPointsIncreaseRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanCrmCustomerPointsIncreaseRequest) CheckParam() error {
	if y.User.AccountType == 0 && y.User.AccountID == "" {
		return errors.New("请指定账号类型和账号ID")
	}
	if y.Points <= 0 {
		return errors.New("points必须大于0的整数")
	}
	if y.Reason == "" {
		return errors.New("积分变更理由必传")
	}
	if y.ClientID == "" {
		return errors.New("三方应用ID必传")
	}
	return nil
}

func (y *YouzanCrmCustomerPointsIncreaseRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
