package domain

type WechatPaymchPackSetting struct {
	ColorName  string `json:"color_name"`
	ColorValue string `json:"color_value"`
	Title      string `json:"title"`
}
type Coupons struct {
	IsAtLeast               int                     `json:"is_at_least"`
	IsWechatPaymchSync      bool                    `json:"is_wechat_paymch_sync"`
	IsSyncWeixin            int                     `json:"is_sync_weixin"`
	EndAt                   int64                   `json:"end_at"`
	IsRandom                int                     `json:"is_random"`
	IsForbidPreference      int                     `json:"is_forbid_preference"`
	Description             string                  `json:"description"`
	Title                   string                  `json:"title"`
	StartAt                 int64                   `json:"start_at"`
	Total                   int                     `json:"total"`
	ExpireNotice            int                     `json:"expire_notice"`
	Quota                   int                     `json:"quota"`
	Stock                   int                     `json:"stock"`
	Value                   string                  `json:"value"`
	StatFetchNum            int                     `json:"stat_fetch_num"`
	ValueRandomTo           string                  `json:"value_random_to"`
	IsShare                 int                     `json:"is_share"`
	StatUseNum              int                     `json:"stat_use_num"`
	StatFetchUserNum        int                     `json:"stat_fetch_user_num"`
	Created                 int64                   `json:"created"`
	RangeType               string                  `json:"range_type"`
	FetchURL                string                  `json:"fetch_url"`
	UserLevelName           string                  `json:"user_level_name"`
	GroupID                 int                     `json:"group_id"`
	WeixinCardID            string                  `json:"weixin_card_id"`
	CouponType              string                  `json:"coupon_type"`
	NeedUserLevel           int                     `json:"need_user_level"`
	AtLeast                 string                  `json:"at_least"`
	Status                  int                     `json:"status"`
	WechatPaymchPackSetting WechatPaymchPackSetting `json:"wechat_paymch_pack_setting,omitempty"`
}
type CouponUnfinished struct {
	Coupons []Coupons `json:"coupons"`
}
