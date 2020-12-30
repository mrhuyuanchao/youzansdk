package domain

type ShopBasicInfo struct {
	// ParentKdtID 上级店铺id，有赞平台生成，在有赞平台唯一，用于判断上一级店铺id
	ParentKdtID int64 `json:"parent_kdt_id"`
	// RootKdtID 有赞连锁总部店铺id，仅供有赞连锁场景下使用。有赞平台生成，在有赞平台唯一，用于判断信息属于哪一个总部
	RootKdtID int64 `json:"root_kdt_id"`
	// CertType 认证类型（0 未认证，2 企业认证，3 个人认证， 12 个体工商户认证， 10 其他组织、政府及事业单位）
	CertType int16 `json:"cert_type"`
	// Name 店铺名称
	Name string `json:"name"`
	// Logo 店铺LOGO
	Logo string `json:"logo"`
	// PhysicalUrl 门店地址
	PhysicalUrl string `json:"physical_url"`
	// Intro 门店简介
	Intro string `json:"intro"`
	// Url 店铺H5地址
	Url string `json:"url"`
	// Sid KdtId 有赞店铺ID
	Sid int64 `json:"sid"`
	// LineOfBusiness 店铺业务线 10000 微商城，20000 教育， 50000 餐饮，60000 美业， 70000 零售， 140000 供货
	LineOfBusiness int32 `json:"line_of_business"`
	// Role 组织角色 0 单店， 1 总部， 2001 门店，2002 网店，3 独立仓，4 合伙人，6 分销供货店铺，7 前置仓
	Role int32 `json:"role"`
}
