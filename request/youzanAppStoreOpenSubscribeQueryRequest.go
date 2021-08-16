package request

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// YouzanAppStoreOpenSubscribeQueryRequest 查询应用订购和授权记录
type YouzanAppStoreOpenSubscribeQueryRequest struct {
	OrderNO string
	// StartTime 时间戳秒
	StartTime int64
	// EndTime 时间戳秒
	EndTime  int64
	PageNO   int32
	PageSize int8
}

var cstZone = time.FixedZone("CST", 8*3600) // 东八
func (y *YouzanAppStoreOpenSubscribeQueryRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanAppStoreOpenSubscribeQueryRequest) GetApiName() string {
	return "youzan.appstore.open.subscribe.query"
}
func (y *YouzanAppStoreOpenSubscribeQueryRequest) GetApiVersion() string {
	return "1.0.2"
}
func (y *YouzanAppStoreOpenSubscribeQueryRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanAppStoreOpenSubscribeQueryRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.OrderNO != "" {
		param["tid"] = y.OrderNO
	}
	if y.PageNO > 0 {
		param["page_no"] = y.PageNO
	} else {
		param["page_no"] = 1
	}
	if y.PageSize > 0 {
		param["page_size"] = y.PageSize
	} else {
		param["page_size"] = 30
	}
	if y.StartTime > 0 {
		param["start_time"] = time.Unix(y.StartTime, 0).In(cstZone).Format("2006-01-02 15:04:05")
	}
	if y.EndTime > 0 {
		param["end_time"] = time.Unix(y.EndTime, 0).In(cstZone).Format("2006-01-02 15:04:05")
	}
	return param
}

func (y *YouzanAppStoreOpenSubscribeQueryRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanAppStoreOpenSubscribeQueryRequest) CheckParam() error {
	if y.PageSize > 100 {
		return errors.New("每页数据最大不能超过100")
	}
	return nil
}

func (y *YouzanAppStoreOpenSubscribeQueryRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
