package response

import "github.com/lscgzwd/youzansdk/domain"

type YouzanScrmCustomerGetResponse struct {
	baseResponse
	Data domain.CustomerGetInfo `json:"data"`
}
