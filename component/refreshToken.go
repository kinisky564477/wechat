package component

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/astaxie/beego"
)

// RefreshToken 刷新令牌
func (t *WxClient) RefreshToken() time.Duration {
	var reTrySec int64 = 60
	token, refreshToken, expire, err := t.getToken()
	if err != nil {
		log.Error("获取 Component-Access-Token 失败", err.Error())
		expire = reTrySec
	}

	// 不允许连续不断调用
	if expire == 0 {
		expire = reTrySec
	}

	if token != "" {
		t.authorizerAccessToken = token
		t.authorizerRefreshToken = refreshToken
	}

	if t.reflashToken != nil && refreshToken != "" {
		var c = map[string]interface{}{
			"token":        token,
			"refreshToken": refreshToken,
			"appid":        t.certificate["appid"],
			"expire_in":    expire,
		}
		t.reflashToken(c)
	}

	return time.Duration(expire) * time.Second
}

// WxTokenParams 刷新令牌参数
type WxTokenParams struct {
	ComponentAppid         string `json:"component_appid"`
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

func (t *WxClient) getToken() (string, string, int64, error) {
	api := API["refresh_access_token"]["post"]
	params := url.Values{}

	params.Set("component_access_token", t.getComponentToken())

	p := WxTokenParams{
		ComponentAppid:         t.getComponentCertificate()["appid"],
		AuthorizerAppid:        t.certificate["appid"],
		AuthorizerRefreshToken: t.authorizerRefreshToken,
	}

	d, err := json.Marshal(p)
	if err != nil {
		beego.Error("转换获取token参数失败,", err)
		return "", "", 0, err
	}

	res, err := t.request.Do(api, params, bytes.NewBuffer(d))
	if err != nil {
		beego.Error("刷新失败：", err)
		return "", "", 0, err
	}

	beego.Error("刷新token返回值:", res)
	token, _ := res["authorizer_access_token"].(string)
	refreshToken, _ := res["authorizer_refresh_token"].(string)
	expire, _ := res["expires_in"].(float64)
	return token, refreshToken, int64(expire), nil
}
