package response

type YouzanScrmCustomerGetResponse struct {
	BaseResponse
	Response CustomerGetInfo `json:"response"`
}

type CustomerGetInfo struct {
	Name     string         `json:"name"`
	Gender   int            `json:"gender"`
	Birthday string         `json:"birthday"`
	Mobile   string         `json:"mobile"`
	IsMember bool           `json:"is_member"`
	Address  ContactAddress `json:"contact_address"`
}
type ContactAddress struct {
	Address  string `json:"address"`
	AreaCode string `json:"areaCode"`
	Province string `jsoon:"province"`
	City     string `json:"city"`
	Country  string `json:"country"`
}
