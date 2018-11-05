package core

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Request 可以直接使用的 http request
type Request interface {
	// Do 自动识别 http method, 请求数据并返回
	Do(API, url.Values, ...io.Reader) (map[string]interface{}, error)
}

// CheckJSONResult 验证 Json 结果
type CheckJSONResult func(map[string]interface{}) error

// DefaultRequest 简易 http request
type DefaultRequest struct {
	checkJSONResult CheckJSONResult
}

// NewDefaultRequest 初始化
func NewDefaultRequest(checkJSONResult CheckJSONResult) *DefaultRequest {
	return &DefaultRequest{
		checkJSONResult: checkJSONResult,
	}
}

// Do 请求远程数据
func (t *DefaultRequest) Do(api API, params url.Values, body ...io.Reader) (map[string]interface{}, error) {
	var req *http.Request
	var err error
	client := &http.Client{}
	URL := api.URI

	if params != nil {
		apiURL, _ := url.Parse(api.URI)
		query := apiURL.Query()
		for k := range params {
			query.Set(k, params.Get(k))
		}
		apiURL.RawQuery = query.Encode()
		URL = apiURL.String()
	}

	log.Info("请求远程接口: ", URL)

	if api.Method == http.MethodGet || len(body) == 0 {
		req, err = http.NewRequest(api.Method, URL, nil)
	} else {
		req, err = http.NewRequest(api.Method, URL, body[0])
	}
	if err != nil {
		log.Error("初始化请求客户端失败", err.Error())
		return nil, err
	}

	if api.ResponseType == ResponseXML {
		req.Header.Add("Content-type", "text/xml")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error("请求远程数据失败", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("解析请求结果失败, ", err.Error())
		return nil, err
	}

	log.Info("远程请求结果:", string(respBody))

	if api.ResponseType == ResponseJSON {
		var res map[string]interface{}
		res, err = t.jsonResult(respBody)
		if err != nil {
			return nil, err
		}

		return res, t.checkJSONResult(res)
	} else if api.ResponseType == ResponseXML {
		var res map[string]interface{}
		res, err = t.xmlResult(respBody)
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, nil
}

func (t *DefaultRequest) jsonResult(body []byte) (map[string]interface{}, error) {
	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		log.Error("转换请求结果(JSON)失败", err.Error())
		return nil, err
	}

	return res, nil
}

func (t *DefaultRequest) xmlResult(body []byte) (map[string]interface{}, error) {
	var res map[string]interface{}
	if err := xml.Unmarshal(body, &res); err != nil {
		log.Error("转换请求结果(XML)失败", err.Error())
		return nil, err
	}

	return res, nil
}

var _ Request = &DefaultRequest{}
