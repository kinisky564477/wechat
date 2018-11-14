package component

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/astaxie/beego"
)

// ComponentAccessTokenTask 刷新令牌
func (t *ComponentClient) ComponentAccessTokenTask() time.Duration {
	var reTrySec int64 = 60
	token, expire, err := t.getComponentToken()
	if err != nil {
		log.Error("获取 Component-Access-Token 失败", err.Error())
		expire = reTrySec
	}
	beego.Error("刷新令牌")

	// 不允许连续不断调用
	if expire == 0 {
		expire = reTrySec
	}

	t.componentAccessToken = token
	return time.Duration(expire) * time.Second
}

// TokenParams 刷新令牌参数
type TokenParams struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"tokenParams"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

// getComponentToken 获取 token
func (t *ComponentClient) getComponentToken() (string, int64, error) {
	api := API["component_token"]["post"]
	params := url.Values{}
	beego.Error("获取token:", api)
	p := TokenParams{
		ComponentAppid:        t.certificate["appid"],
		ComponentAppsecret:    t.certificate["secret"],
		ComponentVerifyTicket: t.certificate["componentVerifyTicket"],
	}

	d, err := json.Marshal(p)
	if err != nil {
		log.Error("转换获取token参数失败,", err)
		return "", 0, err
	}

	res, err := t.request.Do(api, params, bytes.NewBuffer(d))
	if err != nil {
		return "", 0, err
	}
	beego.Error("获取token返回值为:", res)
	token, _ := res["component_access_token"].(string)
	expire, _ := res["expires_in"].(float64)
	return token, int64(expire), nil
}
