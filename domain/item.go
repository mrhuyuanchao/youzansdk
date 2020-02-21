package domain

// Item 商品
type Item struct {
	ItemID           int64                `json:"item_id"`
	Alias            string               `json:"alias"`
	Title            string               `json:"title"`
	Price            int64                `json:"price"`
	ItemType         int64                `json:"item_type"`
	ItemNo           string               `json:"item_no"`
	Quantity         int                  `json:"quantity"`
	PostType         int                  `json:"post_type"`
	CreatedTime      string               `json:"created_time"`
	UpdateTime       string               `json:"update_time"`
	DetailURL        string               `json:"detail_url"`
	Num              int64                `json:"num"`
	Origin           string               `json:"origin"`
	ClassID          string               `json:"classId"`
	Image            string               `json:"image"`
	ShareIcon        string               `json:"shareIcon"`
	ShareTitle       string               `json:"shareTitle"`
	ShareDetail      int64                `json:"shareDetail"`
	PageURL          string               `json:"pageUrl"`
	ItemImgs         []ItemImg            `json:"item_imgs"`
	DeliveryTemplate DeliveryTemplateInfo `json:"delivery_template"`
}
