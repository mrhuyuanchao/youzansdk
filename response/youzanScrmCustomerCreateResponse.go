package response

import "github.com/mrhuyuanchao/youzansdk/domain"

type YouzanScrmCustomerCreateResponse struct {
	baseResponse
	Data domain.CustomerCreateInfo `json:"data"`
}
