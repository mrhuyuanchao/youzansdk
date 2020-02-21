package domain

type PresaleExtend struct {
	PresaleEnd string `json:"presale_end"`
	EtdEnd     string `json:"etd_end"`
	EtdDays    int    `json:"etd_days"`
	EtdStart   string `json:"etd_start"`
	EtdType    int    `json:"etd_type"`
}
