package request

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

type YouzanUmpCouponTakeRequest struct {
	CouponGroupID int64
	FansID        string
	Mobile        string
	OpenUserID    string
	WeixinOpenID  string
}

func (y YouzanUmpCouponTakeRequest) GetMethod() string {
	return "POST"
}
func (y YouzanUmpCouponTakeRequest) GetApiName() string {
	return "youzan.ump.coupon.take"
}
func (y YouzanUmpCouponTakeRequest) GetApiVersion() string {
	return "3.0.0"
}
func (y YouzanUmpCouponTakeRequest) GetParam() map[string]string {
	return nil
}
func (y YouzanUmpCouponTakeRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	param["coupon_group_id"] = strconv.FormatInt(y.CouponGroupID, 10)
	if y.FansID != "" {
		param["fans_id"] = y.FansID
	}
	if y.Mobile != "" {
		param["mobile"] = y.Mobile
	}
	if y.OpenUserID != "" {
		param["open_user_id"] = y.OpenUserID
	}
	if y.WeixinOpenID != "" {
		param["weixin_openid"] = y.WeixinOpenID
	}
	return param
}

func (y YouzanUmpCouponTakeRequest) CheckParam() error {
	if y.CouponGroupID == 0 {
		return errors.New("请设置优惠券ID")
	}
	if y.FansID == "" && y.Mobile == "" && y.OpenUserID == "" && y.WeixinOpenID == "" {
		return errors.New("请设置发送对象")
	}
	return nil
}
func (y YouzanUmpCouponTakeRequest) GetUrlValues() url.Values {
	return nil
}

func (y YouzanUmpCouponTakeRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
