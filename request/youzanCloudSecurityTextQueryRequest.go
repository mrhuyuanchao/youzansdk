package request

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// YouzanCloudSecurityTextQueryRequest 文本非法信息检测
type YouzanCloudSecurityTextQueryRequest struct {
	Content   string
	RequestID string
}

func (y *YouzanCloudSecurityTextQueryRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanCloudSecurityTextQueryRequest) GetApiName() string {
	return "youzan.crm.customer.points.decrease"
}
func (y *YouzanCloudSecurityTextQueryRequest) GetApiVersion() string {
	return "1.0.0"
}
func (y *YouzanCloudSecurityTextQueryRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanCloudSecurityTextQueryRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.RequestID != "" {
		param["data_id"] = y.RequestID
	} else {
		param["data_id"] = time.Now().UnixNano()
	}
	param["content"] = y.Content
	return param
}

func (y *YouzanCloudSecurityTextQueryRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanCloudSecurityTextQueryRequest) CheckParam() error {
	if y.Content == "" {
		return errors.New("请输入需要检测的文本")
	}
	return nil
}

func (y *YouzanCloudSecurityTextQueryRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
