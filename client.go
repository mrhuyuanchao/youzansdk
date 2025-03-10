package youzansdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mrhuyuanchao/youzansdk/request"
	"github.com/mrhuyuanchao/youzansdk/utils"
)

// YouzanClient ..
type YouzanClient struct {
	Token string
	URL   string
}

// Execute 执行请求
func (c YouzanClient) Execute(request request.BaseRequest, v interface{}) (string, error) {
	if err := request.CheckParam(); err != nil {
		return "", err
	}
	// apiName := request.GetApiName()
	// idx := strings.LastIndex(apiName, ".")
	// service := apiName[:idx]
	// action := apiName[idx+1 : len(apiName)]
	url := fmt.Sprintf("%s/%s/%s", c.URL, request.GetApiName(), request.GetApiVersion())
	var response *http.Response
	var err error
	if request.GetMethod() != "POSTFORM" {
		param := request.GetParam()
		param["access_token"] = c.Token
		response, err = utils.HTTP(url, request.GetMethod(), request.GetBodyParam(), param, nil)
	} else {
		url = fmt.Sprintf("%s?access_token=%s", url, c.Token)
		response, err = utils.HttpPostForm(url, request.GetUrlValues())
	}
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("read response body fail")
	}
	err = json.Unmarshal(responseBody, v)

	return string(responseBody), err
}
