package domain

// CouponTake ..
type CouponTake struct {
	CouponType string    `json:"coupon_type"`
	Code       Promocode `json:"promocode"`
	Card       Promocard `json:"promocard"`
}
