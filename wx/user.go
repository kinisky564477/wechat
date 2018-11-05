package wx

import (
	"net/url"
)

// GetUserDetail 获取用户详情(UUID)
func (t *Client) GetUserDetail(openID string) (map[string]interface{}, error) {
	log.Info("获取用户详情(UUID): openID=" + openID)

	api := API["user"]["detail"]
	params := url.Values{}
	params.Set("access_token", t.accessToken)
	params.Set("openid", openID)

	res, err := t.request.Do(api, params)
	if err != nil {
		log.Error("获取用户信息(UUID) 失败, ", err.Error())
		return nil, err
	}

	return res, nil
}
