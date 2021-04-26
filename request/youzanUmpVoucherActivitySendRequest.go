package request

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

type YouzanUmpVoucherActivitySendRequest struct {
	ActivityID   int64
	FansID       string
	Mobile       string
	OpenUserID   string
	WeixinOpenID string
	PlatformUserID string
	VerifyCode string
}

func (y *YouzanUmpVoucherActivitySendRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanUmpVoucherActivitySendRequest) GetApiName() string {
	return "youzan.ump.voucheractivity.send"
}
func (y *YouzanUmpVoucherActivitySendRequest) GetApiVersion() string {
	return "3.0.1"
}
func (y *YouzanUmpVoucherActivitySendRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanUmpVoucherActivitySendRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	param["activity_id"] = strconv.FormatInt(y.ActivityID, 10)
	if y.FansID != "" {
		param["fans_id"] = y.FansID
	}
	if y.Mobile != "" {
		param["mobile"] = y.Mobile
	}
	if y.OpenUserID != "" {
		param["yz_open_id"] = y.OpenUserID
	}
	if y.WeixinOpenID != "" {
		param["weixin_openid"] = y.WeixinOpenID
	}
	if y.PlatformUserID != "" {
		param["platform_user_id"] = y.PlatformUserID
	}
	if y.VerifyCode != "" {
		param["verify_code"] = y.VerifyCode
	}
	return param
}

func (y *YouzanUmpVoucherActivitySendRequest) CheckParam() error {
	if y.ActivityID == 0 {
		return errors.New("请设置优惠券ID")
	}
	if y.FansID == "" && y.Mobile == "" && y.OpenUserID == "" && y.WeixinOpenID == "" {
		return errors.New("请设置发送对象")
	}
	return nil
}
func (y *YouzanUmpVoucherActivitySendRequest) GetUrlValues() url.Values {
	return nil
}

func (y *YouzanUmpVoucherActivitySendRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
