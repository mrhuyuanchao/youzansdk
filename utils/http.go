package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// HTTP 执行HTTP请求
func HTTP(requestUrl, method string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
	var bodyJSON []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJSON, err = json.Marshal(body)
		if err != nil {
			return nil, errors.New("http post body to json fail")
		}
		fmt.Println(string(bodyJSON))
	}
	fmt.Println(requestUrl)
	req, err := http.NewRequest(method, requestUrl, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, errors.New("new request is fail")
	}
	if strings.ToUpper(method) == "POST" {
		req.Header.Set("Content-type", "application/json")
	}
	q := req.URL.Query()
	if params != nil && len(params) > 0 {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	if headers != nil && len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	client := &http.Client{}
	return client.Do(req)
}

// HttpPostForm ...
func HttpPostForm(requestURL string, urlValue url.Values) (*http.Response, error) {
	body := bytes.NewBufferString(urlValue.Encode())
	rsp, err := http.Post(requestURL, "application/x-www-form-urlencoded", body)
	return rsp, err
}
