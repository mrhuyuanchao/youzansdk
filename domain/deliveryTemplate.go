package domain

// DeliveryTemplateInfo 运费模板
type DeliveryTemplateInfo struct {
	DeliveryTemplateID            int64  `json:"delivery_template_id"`
	DeliveryTemplateFee           string `json:"delivery_template_fee"`
	DeliveryTemplateName          string `json:"delivery_template_name"`
	DeliveryTemplateValuationType int    `json:"delivery_template_valuationType"`
}
