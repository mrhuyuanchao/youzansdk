package response

import "github.com/mrhuyuanchao/youzansdk/domain"

type YouzanItemGetResponse struct {
	baseResponse
	Data struct {
		Item struct {
			PostFee              int64                       `json:"post_fee"`
			Skus                 []domain.Sku                `json:"skus"`
			PurchaseRight        bool                        `json:"purchase_right"`
			SkuImages            []domain.SkuImage           `json:"sku_images"`
			PurchaseRightList    domain.PurchaseRight        `json:"purchase_right_list"`
			Desc                 string                      `json:"desc"`
			Price                int64                       `json:"price"`
			DeliveryTemplateInfo domain.DeliveryTemplateInfo `json:"delivery_template_info"`
			ItemImgs             []domain.ItemImg            `json:"item_imgs"`
			Title                string                      `json:"title"`
			PicUrl               string                      `json:"pic_url"`
			Quantity             int64                       `json:"quantity"`
			VirtualExtend        domain.VirtualExtend        `json:"virtual_extend"`
			SellPoint            string                      `json:"sell_point"`
			Created              string                      `json:"created"`
			PicThumbURL          string                      `json:"pic_thumb_url"`
			OriginPrice          string                      `json:"origin_price"`
			IsListing            bool                        `json:"is_listing"`
			AutoListingTime      string                      `json:"auto_listing_time"`
			Template             domain.Template             `json:"template"`
			ItemType             int                         `json:"item_type"`
			TagIds               []int64                     `json:"tag_ids"`
			KdtID                int64                       `json:"kdt_id"`
			JoinLevelDiscount    bool                        `json:"join_level_discount"`
			ItemID               int64                       `json:"item_id"`
			DetailURL            string                      `json:"detail_url"`
			Message              string                      `json:"message"`
			HotelExtend          domain.HotelExtend          `json:"hotel_extend"`
			OneItemMultiCode     string                      `json:"one_item_multi_code"`
			Summary              string                      `json:"summary"`
			ItemTags             []domain.ItemTag            `json:"item_tags"`
			CID                  int64                       `json:"cid"`
			ShareURL             string                      `json:"share_url"`
			PresaleExtend        domain.PresaleExtend        `json:"presale_extend"`
			ItemNo               string                      `json:"item_no"`
			Num                  int64                       `json:"num"`
			SoldNum              int64                       `json:"sold_num"`
			FenxiaoExtend        domain.FenxiaoExtend        `json:"fenxiao_extend"`
			PostType             int                         `json:"post_type"`
			IsLock               bool                        `json:"is_lock"`
			BuyQuota             int64                       `json:"buy_quota"`
			Alias                string                      `json:"alias"`
		} `json:"item"`
	} `json:"data"`
}
