package domain

// ItemTag ..
type ItemTag struct {
	ShareURL string `json:"share_url"`
	Created  string `json:"created"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	ID       int64  `json:"id"`
	TagURL   string `json:"tag_url"`
	Type     int    `json:"type"`
	ItemNum  int    `json:"item_num"`
	Desc     string `json:"desc"`
}
