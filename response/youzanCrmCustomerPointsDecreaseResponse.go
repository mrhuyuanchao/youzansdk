package response

// YouzanCrmCustomerPointsDecreaseResponse 扣除用户的积分
type YouzanCrmCustomerPointsDecreaseResponse struct {
	baseResponse
	Data struct {
		IsSuccess bool `json:"is_success"`
	} `json:"data"`
}
