package response

type YouzanUmpCouponTakeResponse struct {
	BaseResponse
	Response CouponTake `json:"response"`
}

type CouponTake struct {
	CouponType string    `json:"coupon_type"`
	Code       Promocode `json:"promocode"`
	Card       Promocard `json:"promocard"`
}

type Promocode struct {
	PromocodeID     int64  `json:"promocode_id"`
	Title           string `json:"title"`
	Code            string `json:"code"`
	Value           string `json:"value"`
	Condition       string `json:"condition"`
	StartAt         string `json:"start_at"`
	EndAt           string `json:"end_at"`
	IsUsed          int    `json:"is_used"`
	IsInvalid       int    `json:"is_invalid"`
	IsExpired       int    `json:"is_expired"`
	BackgroundColor string `json:"background_color"`
	DetailURL       string `json:"detail_url"`
	VerifyCode      string `json:"verify_code"`
}

type Promocard struct {
	PromocardID     int64  `json:"promocard_id"`
	Title           string `json:"title"`
	Value           string `json:"value"`
	Condition       string `json:"condition"`
	StartAt         string `json:"start_at"`
	EndAt           string `json:"end_at"`
	IsUsed          int    `json:"is_used"`
	IsInvalid       int    `json:"is_invalid"`
	IsExpired       int    `json:"is_expired"`
	BackgroundColor string `json:"background_color"`
	DetailURL       string `json:"detail_url"`
	VerifyCode      string `json:"verify_code"`
}
