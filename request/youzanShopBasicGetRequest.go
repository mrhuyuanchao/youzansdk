package request

import (
	"fmt"
	"net/url"
)

// YouzanShopBasicGetRequest 获取店铺基本信息
type YouzanShopBasicGetRequest struct {
	ItemID int64
	Alias  string
}

func (y YouzanShopBasicGetRequest) GetMethod() string {
	return "POST"
}
func (y YouzanShopBasicGetRequest) GetApiName() string {
	return "youzan.shop.basic.get"
}
func (y YouzanShopBasicGetRequest) GetApiVersion() string {
	return "3.0.0"
}
func (y YouzanShopBasicGetRequest) GetParam() map[string]string {
	return nil
}
// GetBodyParam 组装请求POST参数
func (y YouzanShopBasicGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	return param
}

func (y YouzanShopBasicGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}
// CheckParam 检查参数正确性
func (y YouzanShopBasicGetRequest) CheckParam() error {

	return nil
}
// GetRequestUrl 获取请求URL
func (y YouzanShopBasicGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
