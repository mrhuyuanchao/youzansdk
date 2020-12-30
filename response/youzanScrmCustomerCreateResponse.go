package response

import "github.com/lscgzwd/youzansdk/domain"

type YouzanScrmCustomerCreateResponse struct {
	baseResponse
	Data domain.CustomerCreateInfo `json:"data"`
}
