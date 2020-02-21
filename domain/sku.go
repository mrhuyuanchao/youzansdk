package domain

type Sku struct {
	SkuUniqueCode      string `json:"sku_unique_code"`
	OuterID            string `json:"outer_id"`
	SkuID              int64  `json:"sku_id"`
	OneItemMultiCode   string `json:"one_item_multi_code"`
	SoldNum            int64  `json:"sold_num"`
	ItemNo             string `json:"item_no"`
	Modified           string `json:"modified"`
	CostPrice          int64  `json:"cost_price"`
	PropertiesNameJson string `json:"properties_name_json"`
	Created            string `json:"created"`
	WithHoldQuantity   int64  `json:"with_hold_quantity"`
	Quantity           int64  `json:"quantity"`
	Price              int64  `json:"price"`
	ItemID             int64  `json:"item_id"`
}
