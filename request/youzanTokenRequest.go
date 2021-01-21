package request

import (
	"errors"
	"fmt"
	"net/url"
)

// YouzanTokenRequest ...
type YouzanTokenRequest struct {
	// 'client_id' => $this->clientId,
	//            'client_secret' => $this->clientSecret,
	// 'authorize_type' => 'silent',
	//            'grant_id' => $authorityId,
	//            'grant_type' => 'youzanke',
	//            'refresh' => array_key_exists('refresh', $config) ? boolval($config['refresh']) : false,
	// $params['refresh_token'] = $keys['refresh_token'];
	// $params['code'] = $keys['code'];
	ClientID      string
	ClientSecret  string
	AuthorizeType string
	GrantID       string
	GrantType     string
	Refresh       bool
	RefreshToken  string
	Code          string
}

func (y *YouzanTokenRequest) GetMethod() string {
	return "POST"
}
func (y *YouzanTokenRequest) GetApiName() string {
	return ""
}
func (y *YouzanTokenRequest) GetApiVersion() string {
	return ""
}
func (y *YouzanTokenRequest) GetParam() map[string]string {
	return map[string]string{}
}
func (y *YouzanTokenRequest) GetBodyParam() interface{} {
	param := make(map[string]interface{}, 0)
	param["client_id"] = y.ClientID
	param["client_secret"] = y.ClientSecret
	switch y.AuthorizeType {
	case "authorization_code":
		param["authorize_type"] = y.AuthorizeType
		param["code"] = y.Code
		break
	case "refresh_token":
		param["authorize_type"] = y.AuthorizeType
		param["refresh_token"] = y.RefreshToken
		break
	case "silent":
		param["authorize_type"] = y.AuthorizeType
		param["grant_id"] = y.GrantID
		param["refresh"] = y.Refresh
		break
	case "certificate":
		param["authorize_type"] = y.AuthorizeType
		param["grant_id"] = y.GrantID
		param["refresh"] = y.Refresh
		break
	case "youzanke":
		param["silent"] = y.AuthorizeType
		param["grant_type"] = "youzanke"
		param["grant_id"] = y.GrantID
		param["refresh"] = y.Refresh
		break
	default:
	}
	return param
}

func (y *YouzanTokenRequest) GetUrlValues() url.Values {
	u := url.Values{}
	return u
}

func (y *YouzanTokenRequest) CheckParam() error {
	if y.ClientID == "" && y.ClientSecret == "" {
		return errors.New("client_id和client_secret不能为空")
	}
	switch y.AuthorizeType {
	case "authorization_code":
		if y.Code == "" {
			return errors.New("code不能为空")
		}
		break
	case "refresh_token":
		if y.RefreshToken == "" {
			return errors.New("refresh_token不能为空")
		}
		break
	case "silent":
		if y.GrantID == "" {
			return errors.New("grant_id不能为空")
		}
		break
	case "certificate":
		if y.GrantID == "" {
			return errors.New("grant_id不能为空")
		}
		break
	case "youzanke":
		if y.GrantID == "" {
			return errors.New("grant_id不能为空")
		}
		break
	default:
		return errors.New("授权类型authorize_type错误")
	}
	return nil
}

func (y *YouzanTokenRequest) GetRequestUrl(token string) string {
	return fmt.Sprintf(apiRequestPathToken, youzanApiBaseUrl)
}
