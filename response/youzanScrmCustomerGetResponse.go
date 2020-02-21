package response

import "github.com/mrhuyuanchao/youzansdk/domain"

type YouzanScrmCustomerGetResponse struct {
	baseResponse
	Data domain.CustomerGetInfo `json:"data"`
}
