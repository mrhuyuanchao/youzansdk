package response

import "github.com/lscgzwd/youzansdk/domain"

// YouzanItemsOnsaleGetResponse 获取出售中的商品列表
type YouzanItemsOnsaleGetResponse struct {
	baseResponse
	Data struct {
		Count int           `json:"count"`
		Items []domain.Item `json:"items"`
	} `json:"data"`
}
