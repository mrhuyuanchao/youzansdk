package response

type YouzanMessageSubscriptionApplyResponse struct {
	baseResponse
	Data struct {
		RequestId string `json:"request_id"`
	} `json:"data"`
}
