package component

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// GetPreAuthCode 获取预授权码
func (t *ComponentClient) GetPreAuthCode() (string, error) {
	api := API["pre_auth_code"]["post"]
	params := url.Values{}
	params.Set("component_access_token", t.componentAccessToken)
	type autoParam struct {
		ComponentAppid string `json:"component_appid"`
	}
	p := autoParam{
		ComponentAppid: t.certificate["appid"],
	}
	d, err := json.Marshal(p)
	if err != nil {
		log.Error("转换预授权码参数失败,", err)
		return "", err
	}
	res, err := t.request.Do(api, params, bytes.NewBuffer(d))
	if err != nil {
		log.Error("获取预授权码失败：", err.Error())
		return "", err
	}
	authcode, _ := res["pre_auth_code"].(string)
	return authcode, nil
}
