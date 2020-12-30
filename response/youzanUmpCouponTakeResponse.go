package response

import "github.com/lscgzwd/youzansdk/domain"

type YouzanUmpCouponTakeResponse struct {
	baseResponse
	Data domain.CouponTake `json:"data"`
}
