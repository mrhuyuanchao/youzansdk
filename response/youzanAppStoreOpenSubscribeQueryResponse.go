package response

import "github.com/lscgzwd/youzansdk/domain"

// YouzanAppStoreOpenSubscribeQueryResponse 根据new_user_token拉信息
type YouzanAppStoreOpenSubscribeQueryResponse struct {
	baseResponse
	Data struct {
		Details struct {
			// InnerItemPurchaseNotifyMessageList inner_item_purchase_notify_message_list
			InnerItemPurchaseNotifyMessageList []domain.SubscribeItem `json:"inner_item_purchase_notify_message_list"`
			// AppSubscribeNotifyMessage app_subscribe_notify_message
			AppSubscribeNotifyMessage []domain.SubscribeItem `json:"app_subscribe_notify_message"`
			// AppAuthNotifyMessage app_auth_notify_message
			AppAuthNotifyMessage []domain.SubscribeItem `json:"app_auth_notify_message"`
		} `json:"details"`
		TotalCount int64 `json:"total_count"`
	} `json:"data"`
}
