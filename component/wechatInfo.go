package component

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/astaxie/beego"
)

// GetWechatInfo 获取微信信息
func (t *WxClient) GetWechatInfo() (map[string]interface{}, error) {
	api := API["wechat"]["getinfo"]
	params := url.Values{}
	params.Set("component_access_token", t.getComponentToken())

	type wechatInfoParams struct {
		ComponentAppid  string `json:"component_appid"`
		AuthorizerAppid string `json:"authorizer_appid"`
	}

	p := wechatInfoParams{
		ComponentAppid:  t.getComponentCertificate()["appid"],
		AuthorizerAppid: t.certificate["appid"],
	}

	d, err := json.Marshal(p)
	if err != nil {
		beego.Error("转换获取微信信息参数失败,", err)
		return nil, err
	}

	res, err := t.request.Do(api, params, bytes.NewBuffer(d))
	if err != nil {
		beego.Error("获取微信信息失败：", err)
		return nil, err
	}

	return res, nil
}
