package component

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// AuthorizerToken 获取授权
func (t *WxClient) AuthorizerToken() error {
	api := API["authorizer_access_token"]["post"]
	params := url.Values{}
	params.Set("component_access_token", t.getComponentToken())
	type authParam struct {
		ComponentAppid    string `json:"component_appid"`
		AuthorizationCode string `json:"authorization_code"`
	}
	p := authParam{
		ComponentAppid:    t.certificate["appid"],
		AuthorizationCode: t.authorizationCode,
	}
	d, err := json.Marshal(p)
	if err != nil {
		log.Error("转换授权码参数失败,", err)
		return err
	}
	res, err := t.request.Do(api, params, bytes.NewBuffer(d))
	if err != nil {
		log.Error("获取授权信息失败：", err.Error())
		return err
	}
	t.authorizerAccessToken = res["authorizer_access_token"].(string)
	t.authorizerRefreshToken = res["authorizer_refresh_token"].(string)
	t.authorizerResult = res
	return nil
}
