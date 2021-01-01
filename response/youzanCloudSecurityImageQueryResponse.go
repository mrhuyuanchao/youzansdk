package response

// YouzanCloudSecurityTextQueryResponse 违法关键词检测
type YouzanCloudSecurityImageQueryResponse struct {
	baseResponse
	Data struct {
		// suggestion
		//java.lang.String
		//检测结果 ACCEPT:通过 REVIEW:人工审核 REJECT:拒绝
		//REJECT
		Suggestion string `json:"suggestion"`
		//label
		//java.lang.String
		//风控检测标签 NORMAL:正常 PORN:色情 POLITICS:涉政 ABUSE:辱骂 BAN:违禁 DRUGS:毒品及工具 OTHER:其他
		//POLITIC
		Label string `json:"label"`
		//score
		//java.lang.Double
		//风控检测分数： 0-100之间（含0-100），分值越大，风险越高
		//10.5
		Score float64
	} `json:"data"`
}
