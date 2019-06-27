package response

type YouzanScrmCustomerCreateResponse struct {
	BaseResponse
	Response CustomerCreateInfo `json:"response"`
}

type CustomerCreateInfo struct {
	AccountID   string `json:"account_id"`
	AccountType string `json:"account_type"`
}
