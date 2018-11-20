package component

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/astaxie/beego"
)

// ComponentAccessTokenTask 刷新令牌
func (t *ComponentClient) ComponentAccessTokenTask() time.Duration {
	var reTrySec int64 = 60
	accessToken, expire, err := t.getComponentToken()
	if err != nil {
		beego.Error("获取 Component-Access-Token 失败", err.Error())
		expire = reTrySec
	} else {
		beego.Error(accessToken)
		t.componentAccessToken = accessToken
		if t.updateToken != nil {
			var c = map[string]interface{}{
				"accessToken": accessToken,
				"expire_in":   expire,
			}
			t.updateToken(c)
		}
	}

	// 不允许连续不断调用
	if expire == 0 {
		expire = reTrySec
	}

	return time.Duration(expire) * time.Second
}

// TokenParams 刷新令牌参数
type TokenParams struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

// getComponentToken 获取 token
func (t *ComponentClient) getComponentToken() (string, int64, error) {
	api := API["component_token"]["post"]
	params := url.Values{}

	beego.Error("ticket:", t.componentVerifyTicket)
	if t.componentVerifyTicket == "" {
		log.Error("请等待ticket返回！")
		return "", 0, errors.New("请等待ticket返回！")
	}

	p := TokenParams{
		ComponentAppid:        t.certificate["appid"],
		ComponentAppsecret:    t.certificate["secret"],
		ComponentVerifyTicket: t.componentVerifyTicket,
	}

	beego.Error("tokenParams:", p)

	d, err := json.Marshal(p)
	if err != nil {
		beego.Error("转换获取token参数失败:", err)
		return "", 0, err
	}
	res, err := t.request.Do(api, params, bytes.NewBuffer(d))
	if err != nil {
		beego.Error("获取token失败:", err)
		return "", 0, err
	}
	token, _ := res["component_access_token"].(string)
	expire, _ := res["expires_in"].(float64)
	return token, int64(expire), nil
}
