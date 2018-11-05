package wx

import (
	"net/url"
)

// GetOpenID 获取用户 OpenID
func (t *Client) GetOpenID(code string) (string, error) {
	log.Info("获取用户 OpenID: code=" + code)

	api := API["oauth2"]["access_token"]
	params := url.Values{}
	params.Set("appid", t.certificate["appid"])
	params.Set("secret", t.certificate["secret"])
	params.Set("code", code)

	res, err := t.request.Do(api, params)
	if err != nil {
		log.Error("获取 Oauth2 Access-Token 失败, ", err.Error())
		return "", err
	}

	return res["openid"].(string), nil
}

// GetUserInfo 获取用户详情
func (t *Client) GetUserInfo(code string) (map[string]interface{}, error) {
	log.Info("获取用户详情: code=" + code)

	// 依据 code 获取 openid, access_token
	res, err := t.getOauthToken(code)
	if err != nil {
		return nil, err
	}

	openID := res["openid"].(string)
	token := res["access_token"].(string)

	// 再依据 openid, access_token 获取详情
	api := API["oauth2"]["userinfo"]
	params := url.Values{}
	params.Set("access_token", token)
	params.Set("openid", openID)
	res, err = t.request.Do(api, params)
	if err != nil {
		log.Error("获取 Oauth2 用户信息 失败, ", err.Error())

		// 即使失败也会返回 openid
		res["openid"] = openID
		return res, err
	}

	return res, nil
}

// getOauthToken 获取 Oauth Token
// 没有对 Token 缓存, 官方声明有  7200s 的生存周期, 实际上是用不到
func (t *Client) getOauthToken(code string) (map[string]interface{}, error) {
	log.Info("获取Oauth Token: code=" + code)

	// 依据 code 获取 openid, access_token
	api := API["oauth2"]["access_token"]
	params := url.Values{}
	params.Set("appid", t.certificate["appid"])
	params.Set("secret", t.certificate["secret"])
	params.Set("code", code)

	res, err := t.request.Do(api, params)
	if err != nil {
		log.Error("获取 Oauth2 Access-Token 失败, ", err.Error())
		return nil, err
	}

	return res, nil
}
