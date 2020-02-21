package request

import (
	"errors"
	"net/url"
)

// YouzanTradeGetRequest ..
type YouzanTradeGetRequest struct {
	Tid string
}

func (y YouzanTradeGetRequest) GetMethod() string {
	return "GET"
}
func (y YouzanTradeGetRequest) GetApiName() string {
	return "youzan.trade.get"
}
func (y YouzanTradeGetRequest) GetApiVersion() string {
	return "4.0.0"
}
func (y YouzanTradeGetRequest) GetParam() map[string]string {
	param := make(map[string]string, 0)
	param["tid"] = y.Tid
	return param
}
func (y YouzanTradeGetRequest) GetBodyParam() interface{} {
	return nil
}

func (y YouzanTradeGetRequest) CheckParam() error {
	if y.Tid == "" {
		return errors.New("tid 不能为空")
	}
	return nil
}
func (y YouzanTradeGetRequest) GetUrlValues() url.Values {
	return nil
}
