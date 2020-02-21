package domain

// ItemImg ...
type ItemImg struct {
	Thumbnail string `json:"thumbnail"`
	Created   string `json:"created"`
	Medium    string `json:"medium"`
	ID        int64  `json:"id"`
	URL       string `json:"url"`
	Combine   string `json:"combine"`
}
