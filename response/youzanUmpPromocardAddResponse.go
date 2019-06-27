package response

// YouzanUmpPromocardAddResponse ..
type YouzanUmpPromocardAddResponse struct {
	BaseResponse
	Response UmpCardAddResponse `json:"response"`
}

// UmpCardAddResponse ..
type UmpCardAddResponse struct {
	Promocard Card `json:"promocard"`
}

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

// MarkTag ..
type MarkTag struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// ItemTag ..
type ItemTag struct {
	ShareURL string `json:"share_url"`
	Created  string `json:"created"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	ID       int64  `json:"id"`
	TagURL   string `json:"tag_url"`
	Type     int    `json:"type"`
	ItemNum  int    `json:"item_num"`
	Desc     string `json:"desc"`
}

// ItemImg ...
type ItemImg struct {
	Thumbnail string `json:"thumbnail"`
	Created   string `json:"created"`
	Medium    string `json:"medium"`
	ID        string `json:"id"`
	URL       string `json:"url"`
	Combine   string `json:"combine"`
}

// SpecifyItem ...
type SpecifyItem struct {
	PromotionCid                  int       `json:"promotion_cid"`
	ItemType                      int       `json:"item_type"`
	BuyQuota                      string    `json:"buy_quota"`
	Num                           int       `json:"num"`
	TemplateTitle                 string    `json:"template_title"`
	HasComponent                  bool      `json:"has_component"`
	ItemWeight                    string    `json:"item_weight"`
	Price                         string    `json:"price"`
	DeliveryTemplateName          string    `json:"delivery_template_name"`
	Order                         int       `json:"order"`
	TagIds                        string    `json:"tag_ids"`
	IsSupplierItem                bool      `json:"is_supplier_item"`
	ItemTags                      []ItemTag `json:"item_tags"`
	Created                       string    `json:"created"`
	ItemValidityStart             int       `json:"item_validity_start"`
	IsListing                     bool      `json:"is_listing"`
	IsUsed                        bool      `json:"is_used"`
	OuterBuyURL                   string    `json:"outer_buy_url"`
	StockLocked                   int       `json:"stock_locked"`
	EffectiveType                 int       `json:"effective_type"`
	ShareURL                      string    `json:"share_url"`
	PicThumbURL                   string    `json:"pic_thumb_url"`
	DeliveryTemplateID            int       `json:"delivery_template_id"`
	IsLock                        string    `json:"is_lock"`
	OriginPrice                   string    `json:"origin_price"`
	PicURL                        string    `json:"pic_url"`
	Cid                           int       `json:"cid"`
	Desc                          string    `json:"desc"`
	IsVirtual                     bool      `json:"is_virtual"`
	DetailURL                     string    `json:"detail_url"`
	AutoListingTime               int       `json:"auto_listing_time"`
	Skus                          []string  `json:"skus"`
	PostFee                       string    `json:"post_fee"`
	DeliveryTemplateValuationType int       `json:"delivery_template_valuation_type"`
	NumIid                        string    `json:"num_iid"`
	Title                         string    `json:"title"`
	JoinLevelDiscount             string    `json:"join_level_discount"`
	OuterID                       string    `json:"outer_id"`
	KdtID                         string    `json:"kdt_id"`
	HolidaysAvailable             int       `json:"holidays_available"`
	Alias                         string    `json:"alias"`
	ItemValidityEnd               int       `json:"item_validity_end"`
	ItemImgs                      []ItemImg `json:"item_imgs"`
	SoldNum                       int       `json:"sold_num"`
	ProductType                   string    `json:"product_type"`
	EffectiveDelayHours           int       `json:"effective_delay_hours"`
	TemplateID                    int       `json:"template_id"`
}
