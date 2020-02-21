package domain

// CustomerGetInfo 客户信息
type CustomerGetInfo struct {
	Name      string         `json:"name"`
	Gender    int            `json:"gender"`
	Birthday  string         `json:"birthday"`
	Mobile    string         `json:"mobile"`
	IsMember  bool           `json:"is_member"`
	Remark    string         `json:"remark"`
	OutUserID string         `json:"out_user_id"`
	Address   ContactAddress `json:"contact_address"`
}
