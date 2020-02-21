package response

// YouzanItemQuantityUpdateResponse 全量或增量方式更新SKU库存
type YouzanItemQuantityUpdateResponse struct {
	baseResponse
	Data struct {
		IsSuccess bool `json:"is_success"`
	} `json:"data"`
}
