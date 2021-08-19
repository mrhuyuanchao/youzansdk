package request

import (
	"fmt"
	"net/url"
)

// YouzanMessageSubscriptionApplyRequest 获取店铺基本信息
type YouzanMessageSubscriptionApplyRequest struct {
	OpenUserID string
	Params     map[string]string
	Page string
}

func (y *YouzanMessageSubscriptionApplyRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanMessageSubscriptionApplyRequest) GetApiName() string {
	return "youzan.message.subscription.apply"
}
func (y *YouzanMessageSubscriptionApplyRequest) GetApiVersion() string {
	return "1.0.0"
}
func (y *YouzanMessageSubscriptionApplyRequest) GetParam() map[string]string {
	return map[string]string{}
}

// GetBodyParam 组装请求POST参数
func (y *YouzanMessageSubscriptionApplyRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	var j = make(map[string]interface{})
	j["template_param_map"] = y.Params
	j["yz_open_id"] = y.OpenUserID
	j["page"] = y.Page
	param["request"] = j
	return param
}

func (y *YouzanMessageSubscriptionApplyRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

// CheckParam 检查参数正确性
func (y *YouzanMessageSubscriptionApplyRequest) CheckParam() error {
	return nil
}

// GetRequestUrl 获取请求URL
func (y *YouzanMessageSubscriptionApplyRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
