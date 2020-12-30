package request

import (
	"fmt"
	"net/url"
	"strconv"
)

// YouzanItemsOnsaleGetRequest 获取出售中的商品列表
type YouzanItemsOnsaleGetRequest struct {
	OrderBy         string
	PageNo          int
	PageSize        int
	Q               string
	TagID           int64
	UpdateTimeEnd   int64
	UpdateTimeStart int64
}

func (y YouzanItemsOnsaleGetRequest) GetMethod() string {
	return "POST"
}
func (y YouzanItemsOnsaleGetRequest) GetApiName() string {
	return "youzan.items.onsale.get"
}
func (y YouzanItemsOnsaleGetRequest) GetApiVersion() string {
	return "3.0.0"
}
func (y YouzanItemsOnsaleGetRequest) GetParam() map[string]string {
	return nil
}
func (y YouzanItemsOnsaleGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.OrderBy != "" {
		param["order_by"] = y.OrderBy
	}
	if y.PageNo != 0 {
		param["page_no"] = strconv.Itoa(y.PageNo)
	}
	if y.PageSize != 0 {
		param["page_size"] = strconv.Itoa(y.PageSize)
	}
	if y.Q != "" {
		param["q"] = y.Q
	}
	if y.TagID != 0 {
		param["tag_id"] = strconv.FormatInt(y.TagID, 10)
	}
	if y.UpdateTimeEnd != 0 {
		param["update_time_end"] = strconv.FormatInt(y.UpdateTimeEnd, 10)
	}
	if y.UpdateTimeStart != 0 {
		param["update_time_start"] = strconv.FormatInt(y.UpdateTimeStart, 10)
	}
	return param
}

func (y YouzanItemsOnsaleGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y YouzanItemsOnsaleGetRequest) CheckParam() error {
	return nil
}

func (y YouzanItemsOnsaleGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
