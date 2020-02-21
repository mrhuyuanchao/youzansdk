package domain

type VirtualExtend struct {
	ItemValidityEnd     string `json:"item_validity_end"`
	OperateVersion      int    `json:"operate_version"`
	WeChatTplID         string `json:"we_chat_tpl_id"`
	HolidaysAvailable   bool   `json:"holidays_available"`
	UseAddress          string `json:"use_address"`
	CardServiceTel      string `json:"card_service_tel"`
	CardServiceTelCode  string `json:"card_service_tel_code"`
	EffectiveType       int    `json:"effective_type"`
	ValidityType        int    `json:"validity_type"`
	CardColor           string `json:"card_color"`
	EffectiveDelayHours int    `json:"effective_delay_hours"`
	ItemValidityDay     int    `json:"item_validity_day"`
	Instructions        string `json:"instructions"`
	UpdateWeChatBag     bool   `json:"update_we_chat_bag"`
	CardTitle           string `json:"card_title"`
	ItemValidityStart   string `json:"item_validity_start"`
}
