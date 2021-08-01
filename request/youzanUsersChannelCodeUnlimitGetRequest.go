package request

import (
	"errors"
	"fmt"
	"net/url"
)

// YouzanUsersChannelCodeUnlimitGetRequest 增加用户积分
type YouzanUsersChannelCodeUnlimitGetRequest struct {
	Page    string
	KdtID     int64
	Scene string
	BusinessType  int
	Width    int
	HyaLine bool
}

func (y *YouzanUsersChannelCodeUnlimitGetRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanUsersChannelCodeUnlimitGetRequest) GetApiName() string {
	return "youzan.users.channel.code.unlimit.get"
}
func (y *YouzanUsersChannelCodeUnlimitGetRequest) GetApiVersion() string {
	return "1.0.0"
}
func (y *YouzanUsersChannelCodeUnlimitGetRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanUsersChannelCodeUnlimitGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	param["page"] = y.Page
	param["kdt_id"] = y.KdtID
	param["business_type"] = y.BusinessType
	param["width"] = y.Width
	param["scene"] = y.Scene
	param["hya_line"] = y.HyaLine
	return map[string]interface{}{"params": param}
}

func (y *YouzanUsersChannelCodeUnlimitGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanUsersChannelCodeUnlimitGetRequest) CheckParam() error {
	if y.KdtID == 0 {
		return errors.New("店铺ID必传")
	}
	if y.Width == 0 {
		y.Width = 400
	}
	if y.BusinessType == 0 {
		y.BusinessType = 1
	}
	if y.Width < 280 || y.Width > 1280 {
		return errors.New("width必须介于280到1280之间")
	}
	if y.Page == "" {
		return errors.New("page页面路径必传")
	}
	if len(y.Scene) > 32 {
		return errors.New("scene长度不能超过32")
	}
	return nil
}

func (y *YouzanUsersChannelCodeUnlimitGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
