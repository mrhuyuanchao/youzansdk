package request

import (
	"errors"
	"fmt"
	"net/url"
)

// YouzanAppStoreOpenTradeCreateRequest 用于创建应用内购订单
type YouzanAppStoreOpenTradeCreateRequest struct {
	KdtID      int64
	UserID     string
	Num        int
	OutItemID  string
	BuyerPhone string
	AppID      int64
	ClientID   string
}

func (y *YouzanAppStoreOpenTradeCreateRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanAppStoreOpenTradeCreateRequest) GetApiName() string {
	return "youzan.appstore.open.trade.create"
}
func (y *YouzanAppStoreOpenTradeCreateRequest) GetApiVersion() string {
	return "1.0.2"
}
func (y *YouzanAppStoreOpenTradeCreateRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanAppStoreOpenTradeCreateRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.KdtID > 0 {
		param["kdt_id"] = y.KdtID
	}
	if y.ClientID != "" {
		param["client_id"] = y.ClientID
	}
	if y.UserID != "" {
		param["user_id"] = y.UserID
	}
	if y.Num > 0 {
		param["num"] = y.Num
	} else {
		param["num"] = 1
	}
	if y.BuyerPhone != "" {
		param["buyer_phone"] = y.BuyerPhone
	}
	param["out_item_id"] = y.OutItemID
	param["app_id"] = y.AppID

	return param
}

func (y *YouzanAppStoreOpenTradeCreateRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanAppStoreOpenTradeCreateRequest) CheckParam() error {
	if y.AppID < 1 {
		return errors.New("app_id应用编号必传")
	}
	if y.OutItemID == "" {
		return errors.New("out_item_id外部商品ID必传")
	}
	if y.UserID == "" && y.BuyerPhone == "" {
		return errors.New("user_id和buyer_phone不能同时为空")
	}
	return nil
}

func (y *YouzanAppStoreOpenTradeCreateRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
