package request

import (
"fmt"
"net/url"
)

// YouzanMessageSubscriptionTemplateGetRequest 获取店铺基本信息
type YouzanMessageSubscriptionTemplateGetRequest struct {
	
}

func (y *YouzanMessageSubscriptionTemplateGetRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanMessageSubscriptionTemplateGetRequest) GetApiName() string {
	return "youzan.message.subscription.template.get"
}
func (y *YouzanMessageSubscriptionTemplateGetRequest) GetApiVersion() string {
	return "1.0.0"
}
func (y *YouzanMessageSubscriptionTemplateGetRequest) GetParam() map[string]string {
	return map[string]string{}
}
// GetBodyParam 组装请求POST参数
func (y *YouzanMessageSubscriptionTemplateGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	param["subscriptionTemplateListGetRequest"] = map[string]string{
		"scene": "yzCloudOpenActivityMessage",
	}
	return param
}

func (y *YouzanMessageSubscriptionTemplateGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}
// CheckParam 检查参数正确性
func (y *YouzanMessageSubscriptionTemplateGetRequest) CheckParam() error {

	return nil
}
// GetRequestUrl 获取请求URL
func (y *YouzanMessageSubscriptionTemplateGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}

