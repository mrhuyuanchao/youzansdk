package domain

// ContactAddress 联系地址
type ContactAddress struct {
	Address  string `json:"address"`
	AreaCode string `json:"areaCode"`
	Province string `jsoon:"province"`
	City     string `json:"city"`
	Country  string `json:"country"`
}
