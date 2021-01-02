package response

// YouzanTokenResponse access_token请求响应
type YouzanTokenResponse struct {
	baseResponse
	Data struct {
		Expires      int64  `json:"expires"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		AuthorityID  string `json:"authority_id"`
	} `json:"data"`
}
