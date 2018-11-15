package component

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/astaxie/beego"
)

// AuthorizerToken 获取授权
func (t *WxClient) AuthorizerToken() error {
	api := API["authorizer_access_token"]["post"]
	params := url.Values{}
	beego.Error(t.getComponentToken())
	params.Set("component_access_token", t.getComponentToken())
	type authParam struct {
		ComponentAppid    string `json:"component_appid"`
		AuthorizationCode string `json:"authorization_code"`
	}

	p := authParam{
		ComponentAppid:    t.getComponentCertificate()["appid"],
		AuthorizationCode: t.authorizationCode,
	}
	d, err := json.Marshal(p)
	if err != nil {
		beego.Error("转换授权码参数失败,", err)
		return err
	}

	res, err := t.request.Do(api, params, bytes.NewBuffer(d))
	if err != nil {
		beego.Error("获取授权信息失败：", err.Error())
		return err
	}
	authorizationInfo := res["authorization_info"].(map[string]interface{})

	t.authorizerAccessToken = authorizationInfo["authorizer_access_token"].(string)
	t.authorizerRefreshToken = authorizationInfo["authorizer_refresh_token"].(string)
	t.authorizerResult = authorizationInfo
	return nil
}
