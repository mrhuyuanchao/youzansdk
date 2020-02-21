package domain

// Promocode 优惠码
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
