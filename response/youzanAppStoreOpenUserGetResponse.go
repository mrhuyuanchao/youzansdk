package response

// YouzanAppStoreOpenUserGetResponse 根据new_user_token拉信息
type YouzanAppStoreOpenUserGetResponse struct {
	baseResponse
	Data struct {
		// shop_role
		//java.lang.Integer
		//店铺角色： 0:单店 1:总店 2:分店
		//1
		ShopRole int8 `json:"shop_role"`
		//shop_type
		//java.lang.Integer
		//店铺类型： 0:微商城 1:微小店 6:美业 7:零售
		//0
		ShopType int8 `json:"shop_type"`
		//phone
		//java.lang.String
		//电话
		//15866666666
		Phone string `json:"phone"`
		//kdt_id
		//java.lang.Long
		//店铺id
		//554134
		KdtID int64 `json:"kdt_id"`
		//user_id
		//java.lang.Long
		//购买者身份标示
		//3413423
		UserID int64 `json:"user_id"`
		//shop_name
		//java.lang.String
		//店铺名称
		//测试店铺
		ShopName string `json:"shop_name"`
		//yz_open_id
		//java.lang.String
		//有赞客户id(推荐使用)
		//4asf1314123
		YzOpenID string `json:"yz_open_id"`
		//app_id
		//java.lang.Long
		//应用id
		//43513
		AppID int64 `json:"app_id"`
		//code
		//java.lang.String
		//授权code
		//13asdfa34131sfadsf234
		Code string `json:"code"`
	} `json:"data"`
}
