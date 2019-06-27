package youzansdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mrhuyuanchao/youzansdk/request"
	"github.com/mrhuyuanchao/youzansdk/utils"
)

// YouzanClient ..
type YouzanClient struct {
	Token string
}

// Execute 执行请求
func (c YouzanClient) Execute(request request.BaseRequest, v interface{}) (string, error) {
	if err := request.CheckParam(); err != nil {
		return "", err
	}
	apiName := request.GetApiName()
	idx := strings.LastIndex(apiName, ".")
	service := apiName[:idx]
	action := apiName[idx+1 : len(apiName)]
	url := fmt.Sprintf("%s/%s/%s/%s", request.GetURL(), service, request.GetApiVersion(), action)
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

// func main() {
// 	// request := request.YouzanUmpPromocardAddRequest{}
// 	// request.AtLeast = 199
// 	// request.DateType = 1
// 	// request.Description = "测试优惠券创建"
// 	// request.Endat = "2019-05-16 23:59:59"
// 	// request.IsAtLeast = 1
// 	// request.NeedUserLevel = 0
// 	// request.Quota = 1
// 	// request.RangeType = "ALL"
// 	// request.StartAt = "2019-05-16 00:00:00"
// 	// request.Title = "会员中心测试优惠券"
// 	// request.Total = 1
// 	// request.Value = 100
// 	// request.PreferentialType = 1
// 	// response := response.YouzanUmpPromocardAddResponse{}

// 	// request := request.YouzanUmpCouponTakeRequest{}
// 	// request.CouponGroupID = 5589603
// 	// //request.WeixinOpenID = "oCquvwgZ67jcj1hNM5G8-KkgRuGE"
// 	// request.Mobile = "15733053717"
// 	// response := response.YouzanUmpCouponTakeResponse{}
// 	// request := request.YouzanScrmCustomerCreateRequest{}
// 	// request.Mobile = "18618439108"
// 	// response := response.YouzanScrmCustomerCreateResponse{}

// 	request := request.YouzanScrmCustomerGetRequest{}
// 	request.AccountID = "15225095001"
// 	request.AccountType = "Mobile"
// 	response := response.YouzanScrmCustomerGetResponse{}
// 	client := YouzanClient{}
// 	client.Token = "5f94afd843293dfa9181210f0fbdca34"
// 	cont, err := client.Execute(request, &response)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(cont)
// 	}
// 	fmt.Println(response)
// }

// access_token=5f94afd843293dfa9181210f0fbdca34&
// at_least=199.00&
// can_give_friend=0&
// date_type=1&
// description=%E6%B5%8B%E8%AF%95%E4%BC%98%E6%83%A0%E5%88%B8%E5%88%9B%E5%BB%BA&
// end_at=2019-05-14+23%3A59%3A59&
// expire_notice=0&
// is_at_least=1&
// is_forbid_preference=0&
// is_random=0&
// is_share=0&
// is_sync_weixin=0&
// need_user_level=0&
// preferential_type=0&
// quota=1&
// range_type=ALL&
// start_at=2019-05-13+14%3A00%3A00&
// title=%E4%BC%9A%E5%91%98%E4%B8%AD%E5%BF%83%E6%B5%8B%E8%AF%95%E4%BC%98%E6%83%A0%E5%88%B8&
// total=1&
// value=100.00
