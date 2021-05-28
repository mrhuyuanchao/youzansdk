package response

type YouzanMessageSubscriptionTemplateGetResponse struct {
	baseResponse
	Data struct {
		TemplateIdList []string `json:"template_id_list"`
	} `json:"data"`
}
