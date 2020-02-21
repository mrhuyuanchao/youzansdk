package domain

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