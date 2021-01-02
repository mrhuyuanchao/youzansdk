package request

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// YouzanItemGetRequest 获取单个商品信息
type YouzanItemGetRequest struct {
	ItemID int64
	Alias  string
}

func (y *YouzanItemGetRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanItemGetRequest) GetApiName() string {
	return "youzan.item.get"
}
func (y *YouzanItemGetRequest) GetApiVersion() string {
	return "3.0.0"
}
func (y *YouzanItemGetRequest) GetParam() map[string]string {
	return nil
}
func (y *YouzanItemGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.ItemID != 0 {
		param["item_id"] = strconv.FormatInt(y.ItemID, 10)
	}
	if y.Alias != "" {
		param["alias"] = y.Alias
	}
	return param
}

func (y *YouzanItemGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanItemGetRequest) CheckParam() error {
	if y.ItemID == 0 && y.Alias == "" {
		return errors.New("item_id和alias不能同时为空")
	}
	return nil
}

func (y *YouzanItemGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
