package response

import "github.com/mrhuyuanchao/youzansdk/domain"

type YouzanUmpCouponTakeResponse struct {
	baseResponse
	Data domain.CouponTake `json:"data"`
}
