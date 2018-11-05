package wx

import (
	"net/url"
	"time"
)

// AccessTokenTask 刷新 任务
func (t *Client) AccessTokenTask() time.Duration {
	var reTrySec int64 = 60
	token, expire, err := t.getToken()
	if err != nil {
		log.Error("获取 Access-Token 失败", err.Error())
		expire = reTrySec
	}

	// 不允许连续不断调用
	if expire == 0 {
		expire = reTrySec
	}

	t.accessToken = token
	return time.Duration(expire) * time.Second
}

// getToken 获取 token
func (t *Client) getToken() (string, int64, error) {
	api := API["access_token"]["get"]
	params := url.Values{}
	params.Set("appid", t.certificate["appid"])
	params.Set("secret", t.certificate["secret"])

	res, err := t.request.Do(api, params)
	if err != nil {
		return "", 0, err
	}

	token, _ := res["access_token"].(string)
	expire, _ := res["expires_in"].(float64)
	return token, int64(expire), nil
}
