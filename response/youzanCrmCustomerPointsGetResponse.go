package response

// YouzanCrmCustomerPointsGetResponse 获取用户的积分
type YouzanCrmCustomerPointsGetResponse struct {
	baseResponse
	Data struct {
		Point int64 `json:"point"`
	} `json:"data"`
}
