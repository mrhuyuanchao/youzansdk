package domain

// Card ...
type Card struct {
	IsAtLeast          int           `json:"is_at_least"`
	WeixinTitle        string        `json:"weixin_title"`
	EndAt              string        `json:"end_at"`
	IsSyncWeixin       int           `json:"is_sync_weixin"`
	DateType           int           `json:"date_type"`
	PreferentialType   int           `json:"preferential_type"`
	IsRandom           int           `json:"is_random"`
	IsForbidPreference int           `json:"is_forbid_preference"`
	MarkTags           []MarkTag     `json:"mark_tags"`
	WeixinColor        string        `json:"weixin_color"`
	Description        string        `json:"description"`
	SpecifyItems       []SpecifyItem `json:"specify_items"`
	Discount           string        `json:"discount"`
	CanGiveFriend      int           `json:"can_give_friend"`
	Title              string        `json:"title"`
	StartAt            string        `json:"start_at"`
	Total              int64         `json:"total"`
	ExpireNotice       int           `json:"expire_notice"`
	Quota              int           `json:"quota"`
	FixedTerm          int           `json:"fixed_term"`
	Stock              int           `json:"stock"`
	Value              string        `json:"value"`
	StatFetchNum       int           `json:"stat_fetch_num"`
	ValueRandomTo      string        `json:"value_random_to"`
	WeixinSubTitle     string        `json:"weixin_sub_title"`
	IsShare            int           `json:"is_share"`
	StatUseNum         int           `json:"stat_use_num"`
	StatFetchUserNum   int           `json:"stat_fetch_user_num"`
	Created            string        `json:"created"`
	WeixinColorRgb     string        `json:"weixin_color_rgb"`
	RangeType          string        `json:"range_type"`
	FetchURL           string        `json:"fetch_url"`
	ServicePhone       string        `json:"service_phone"`
	UserLevelName      string        `json:"user_level_name"`
	IsExpired          int           `json:"is_expired"`
	IsOngoing          int           `json:"is_ongoing"`
	GroupID            int64         `json:"group_id"`
	NeedUserLevel      int           `json:"need_user_level"`
	AtLeast            string        `json:"at_least"`
	Status             int           `json:"status"`
	FixedBeginTerm     int           `json:"fixed_begin_term"`
}
