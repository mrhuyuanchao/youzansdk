package response

// YouzanCrmCustomerPointsIncreaseResponse 增加用户积分
type YouzanCrmCustomerPointsIncreaseResponse struct {
	baseResponse
	Data struct {
		IsSuccess bool `json:"is_success"`
	} `json:"data"`
}
