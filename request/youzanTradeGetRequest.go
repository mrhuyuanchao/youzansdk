package request

import (
	"errors"
	"fmt"
	"net/url"
)

// YouzanTradeGetRequest ..
type YouzanTradeGetRequest struct {
	Tid string
}

func (y *YouzanTradeGetRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanTradeGetRequest) GetApiName() string {
	return "youzan.trade.get"
}
func (y *YouzanTradeGetRequest) GetApiVersion() string {
	return "4.0.0"
}
func (y *YouzanTradeGetRequest) GetParam() map[string]string {
	return nil
}
func (y *YouzanTradeGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	param["tid"] = y.Tid
	return param
}

func (y *YouzanTradeGetRequest) CheckParam() error {
	if y.Tid == "" {
		return errors.New("tid 不能为空")
	}
	return nil
}
func (y *YouzanTradeGetRequest) GetUrlValues() url.Values {
	return nil
}

func (y *YouzanTradeGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
