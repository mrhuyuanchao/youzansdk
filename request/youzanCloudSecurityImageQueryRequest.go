package request

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// YouzanCloudSecurityImageQueryRequest 文本非法信息检测
type YouzanCloudSecurityImageQueryRequest struct {
	ImageUrl  string
	RequestID string
}

func (y *YouzanCloudSecurityImageQueryRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanCloudSecurityImageQueryRequest) GetApiName() string {
	return "youzan.crm.customer.points.decrease"
}
func (y *YouzanCloudSecurityImageQueryRequest) GetApiVersion() string {
	return "1.0.0"
}
func (y *YouzanCloudSecurityImageQueryRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanCloudSecurityImageQueryRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	if y.RequestID != "" {
		param["data_id"] = y.RequestID
	} else {
		param["data_id"] = time.Now().UnixNano()
	}
	param["image_url"] = y.ImageUrl
	return param
}

func (y *YouzanCloudSecurityImageQueryRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanCloudSecurityImageQueryRequest) CheckParam() error {
	if y.ImageUrl == "" {
		return errors.New("请输入需要检测的图片地址")
	}
	return nil
}

func (y *YouzanCloudSecurityImageQueryRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPath, youzanApiBaseUrl, y.GetApiName(), y.GetApiVersion(), token)
}
