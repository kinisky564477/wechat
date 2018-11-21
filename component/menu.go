package component

import (
	"bytes"
	"net/url"
)

// RefreshMenu 刷新菜单
func (t *WxClient) RefreshMenu(menu []byte) error {
	log.Info("准备刷新菜单")
	log.Info("接收到菜单信息:", string(menu))

	api := API["menu"]["delete"]
	params := url.Values{}
	params.Set("access_token", t.authorizerAccessToken)

	_, err := t.request.Do(api, params)
	if err != nil {
		log.Error("清除原有菜单失败", err.Error())
		return err
	}

	if string(menu) != "" {
		api = API["menu"]["create"]
		params = url.Values{}
		params.Set("access_token", t.authorizerAccessToken)

		_, err = t.request.Do(api, params, bytes.NewBuffer(menu))
		if err != nil {
			log.Error("创建菜单失败", err.Error())
			return err
		}

		log.Info("创建菜单成功")
	}
	return nil
}

// GetMenu 获取菜单
func (t *WxClient) GetMenu() (map[string]interface{}, error) {
	log.Info("准备获取菜单")

	api := API["menu"]["get"]
	params := url.Values{}
	params.Set("access_token", t.authorizerAccessToken)

	res, err := t.request.Do(api, params)
	if err != nil {
		log.Error("查询菜单失败", err.Error())
		return nil, err
	}

	return res, nil
}
