package request

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// YouzanUmpPromocardAddRequest ..
type YouzanUmpPromocardAddRequest struct {
	AtLeast            float64
	CanGiveFriend      int
	DateType           int
	Description        string
	Discount           float64
	Endat              string
	ExpireNotice       int
	FixedBeginTerm     int
	FixedTerm          int
	IsAtLeast          int
	IsForbidPerference int
	IsRandom           int
	IsShare            int
	IsSyncWeixin       int
	MarkTagIds         string
	NeedUserLevel      int
	PreferentialType   int
	Quota              int
	RangeType          string
	ServicePhone       string
	SpecifyItemIds     string
	StartAt            string
	Title              string
	Total              int
	Value              float64
	ValueRandomTo      float64
	WeixinColor        string
	WeixinColorRgb     string
	WeixinSubTitle     string
	WeixinTitle        string
}

func (y YouzanUmpPromocardAddRequest) GetURL() string {
	return "https://open.youzan.com/api/oauthentry"
}
func (y YouzanUmpPromocardAddRequest) GetMethod() string {
	return "GET"
}
func (y YouzanUmpPromocardAddRequest) GetApiName() string {
	return "youzan.ump.promocard.add"
}
func (y YouzanUmpPromocardAddRequest) GetApiVersion() string {
	return "3.0.0"
}
func (y YouzanUmpPromocardAddRequest) GetParam() map[string]string {
	param := make(map[string]string, 0)
	param["end_at"] = y.Endat
	param["start_at"] = y.StartAt
	param["title"] = y.Title
	param["is_at_least"] = strconv.Itoa(y.IsAtLeast)
	param["can_give_friend"] = strconv.Itoa(y.CanGiveFriend)
	param["date_type"] = strconv.Itoa(y.DateType)
	param["need_user_level"] = strconv.Itoa(y.NeedUserLevel)
	if y.DateType == 2 {
		param["fixed_begin_term"] = strconv.Itoa(y.FixedBeginTerm)
		param["fixed_term"] = strconv.Itoa(y.FixedTerm)
	}
	if y.IsAtLeast == 1 {
		param["at_least"] = fmt.Sprintf("%.2f", y.AtLeast)
	}
	param["range_type"] = y.RangeType
	if y.RangeType == "PART" {
		param["specify_item_ids"] = y.SpecifyItemIds
	}
	if y.Description != "" {
		param["description"] = y.Description
	}
	if y.Discount > 0 {
		param["discount"] = fmt.Sprintf("%.2f", y.Discount)
	}
	param["expire_notice"] = strconv.Itoa(y.ExpireNotice)
	param["is_forbid_preference"] = strconv.Itoa(y.IsForbidPerference)
	param["is_share"] = strconv.Itoa(y.IsShare)
	if y.MarkTagIds != "" {
		param["mark_tag_ids"] = y.MarkTagIds
	}
	if y.PreferentialType != 0 {
		param["preferential_type"] = strconv.Itoa(y.PreferentialType)
	}
	param["quota"] = strconv.Itoa(y.Quota)
	param["total"] = strconv.Itoa(y.Total)
	param["value"] = fmt.Sprintf("%.2f", y.Value)
	if y.IsRandom == 1 {
		param["is_random"] = strconv.Itoa(y.IsRandom)
		param["value_random_to"] = fmt.Sprintf("%.2f", y.ValueRandomTo)
	}
	if y.IsSyncWeixin == 1 {
		param["is_sync_weixin"] = strconv.Itoa(y.IsSyncWeixin)
		param["weixin_color"] = y.WeixinColor
		param["weixin_color_rgb"] = y.WeixinColorRgb
		param["weixin_sub_title"] = y.WeixinSubTitle
		param["weixin_title"] = y.WeixinTitle
		param["service_phone"] = y.ServicePhone
	}
	return param
}
func (y YouzanUmpPromocardAddRequest) GetBodyParam() interface{} {
	return nil
}

func (y YouzanUmpPromocardAddRequest) CheckParam() error {
	if y.DateType == 0 {
		y.DateType = 1
	}
	if y.PreferentialType != 1 && y.PreferentialType != 2 {
		y.PreferentialType = 1
	}
	if y.Title == "" {
		return errors.New("优惠券标题不能为空")
	}
	if y.Total <= 0 {
		return errors.New("优惠券库存必须大于0")
	}
	if y.StartAt == "" {
		return errors.New("开始时间不能为空")
	}
	if y.Endat == "" {
		return errors.New("结束时间不能为空")
	}
	if y.PreferentialType == 1 && y.IsRandom != 1 && y.Value <= 0 {
		return errors.New("请设置优惠金额")
	}
	if y.PreferentialType == 2 && y.Discount <= 0 {
		return errors.New("请设置优惠折扣")
	}
	if y.RangeType != "PART" && y.RangeType != "ALL" {
		return errors.New("可用范围的类型错误")
	}
	if y.RangeType == "PART" && y.SpecifyItemIds == "" {
		return errors.New("可用范围为PART时商品ID列表不能为空")
	}
	if y.DateType == 2 && (y.FixedTerm <= 0 || y.FixedBeginTerm <= 0) {
		return errors.New("请设置延迟时间")
	}

	return nil
}

func (y YouzanUmpPromocardAddRequest) GetUrlValues() url.Values {
	return nil
}
