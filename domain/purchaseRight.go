package domain

type PurchaseRight struct {
	UmpTagsText  []string `json:"ump_tags_text"`
	UmpTags      []string `json:"ump_tags"`
	UmpLevels    []string `json:"ump_levels"`
	UmpLevelText []string `json:"ump_level_text"`
}
