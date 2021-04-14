package response

import "github.com/lscgzwd/youzansdk/domain"

type YouzanUmpCouponsUnfinishedResponse struct {
	baseResponse
	Data domain.CouponUnfinished `json:"data"`
}
