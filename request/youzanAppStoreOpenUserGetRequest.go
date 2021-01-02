package request

import (
	"errors"
	"fmt"
	"net/url"
)

// YouzanAppStoreOpenUserGetRequest 通过new_user_token获取用户信息
type YouzanAppStoreOpenUserGetRequest struct {
	UserToken string
}

func (y *YouzanAppStoreOpenUserGetRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanAppStoreOpenUserGetRequest) GetApiName() string {
	return "youzan.appstore.open.user.get"
}
func (y *YouzanAppStoreOpenUserGetRequest) GetApiVersion() string {
	return "1.0.0"
}
func (y *YouzanAppStoreOpenUserGetRequest) GetParam() map[string]string {
	return nil
}
func (y *YouzanAppStoreOpenUserGetRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.UserToken != "" {
		param["user_token"] = y.UserToken
	}
	return param
}

func (y *YouzanAppStoreOpenUserGetRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanAppStoreOpenUserGetRequest) CheckParam() error {
	if y.UserToken == "" {
		return errors.New("user_token不能为空")
	}
	return nil
}

func (y *YouzanAppStoreOpenUserGetRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
