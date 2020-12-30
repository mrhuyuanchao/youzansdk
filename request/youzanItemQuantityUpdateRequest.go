package request

import (
	"errors"
	"fmt"
	"net/url"
)

//YouzanItemQuantityUpdateRequest 全量或增量方式更新SKU库存
type YouzanItemQuantityUpdateRequest struct {
	ItemID   int64
	Quantity int
	SkuID    int64
	Type     int
}

func (y YouzanItemQuantityUpdateRequest) GetMethod() string {
	return "POST"
}
func (y YouzanItemQuantityUpdateRequest) GetApiName() string {
	return "youzan.item.quantity.update"
}
func (y YouzanItemQuantityUpdateRequest) GetApiVersion() string {
	return "3.0.0"
}
func (y YouzanItemQuantityUpdateRequest) GetParam() map[string]string {
	return nil
}
func (y YouzanItemQuantityUpdateRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	param["item_id"] = y.ItemID
	param["quantity"] = y.Quantity
	if y.SkuID != 0 {
		param["sku_id"] = y.SkuID
	}
	param["type"] = y.Type
	return param
}

func (y YouzanItemQuantityUpdateRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y YouzanItemQuantityUpdateRequest) CheckParam() error {
	if y.ItemID == 0 {
		return errors.New("item_id不能为空")
	}
	return nil
}

func (y YouzanItemQuantityUpdateRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
