package domain

type SubscribeItem struct {
	// kdt_id
	//java.lang.Long
	//订购店铺id
	//18163424
	KdtID int64 `json:"kdt_id"`
	//effect_time
	//java.lang.Long
	//生效时间， Unix时间戳，单位：毫秒
	//1558949704000
	EffectTime int64 `json:"effect_time"`
	//expire_time
	//java.lang.Long
	//过期时间 ，Unix时间戳，单位：毫秒
	//1558949704000
	ExpireTime int64 `json:"expire_time"`
	//pay_time
	//java.lang.Long
	//支付时间， Unix时间戳，单位：毫秒
	//1558949704000
	PayTime int64 `json:"pay_time"`
	//buyer_id
	//java.lang.Long
	//订购人id
	//865115939
	BuyerID int64 `json:"buyer_id"`
	//app_id
	//java.lang.Long
	//应用id
	//42516
	AppID int64 `json:"app_id"`
	//price
	//java.lang.Long
	//支付价格，单位：分
	//99
	Price int32 `json:"price"`
	//status
	//java.lang.Integer
	//支付状态：10-未支付，20-已支付
	//20
	Status int8 `json:"status"`
	//buyer_phone
	//java.lang.String
	//订购人手机号
	//135xxxx1234
	BuyerPhone string `json:"buyer_phone"`
	//shop_dis_play_no
	//java.lang.String
	//订购店铺编号
	//48567648
	ShopDisPlayNo string `json:"shop_dis_play_no"`
	//out_item_id
	//java.lang.String
	//外部商品编码
	//yz666
	OutItemID string `json:"out_item_id"`
	//order_no
	//java.lang.String
	//订单编号
	//E20190509181910080300061
	OrderNO string `json:"order_no"`
	// SkuIntervalText sku_interval_text
	SkuIntervalText string `json:"sku_interval_text"`
	// SkuVersionText sku_version_text
	SkuVersionText string `json:"sku_version_text"`
}
