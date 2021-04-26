package response

// @see https://doc.youzanyun.com/detail/API/0/979
type YouzanUmpVoucherActivitySendResponse struct {
	baseResponse
	Data struct {
		VoucherIdentity struct {
			CouponID   int64 `json:"coupon_id"`
			CouponType int   `json:"coupon_type"`
		} `json:"voucher_identity"`
		SendSource       string `json:"send_source"`
		ValidStartTime   int64  `json:"valid_start_time"`
		SendKdtID        int    `json:"send_kdt_id"`
		VerifyCode       string `json:"verify_code"`
		SentAt           int64  `json:"sent_at"`
		KdtID            int    `json:"kdt_id"`
		PreferentialMode int    `json:"preferential_mode"`
		ActivityID       int    `json:"activity_id"`
		CodeValue        string `json:"code_value"`
		ValidEndTime     int64  `json:"valid_end_time"`
		Value            int    `json:"value"`
		Status           int    `json:"status"`
	}
}
