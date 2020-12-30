package response

import "github.com/lscgzwd/youzansdk/domain"

// YouzanShopBasicGetResponse 店铺基本信息
type YouzanShopBasicGetResponse struct {
	baseResponse
	Data domain.ShopBasicInfo `json:"data"`
}
