package component

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
)

// GetMaterialList 获取素材列表
func (t *WxClient) GetMaterialList(pars map[string]string) (interface{}, error) {
	api := API["material"]["list"]
	params := url.Values{}
	params.Set("access_token", t.authorizerAccessToken)
	type mParam struct {
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Count  int    `json:"count"`
	}
	var offset = 0
	if pars["offset"] != "" {
		offset, _ = strconv.Atoi(pars["offset"])
	}

	var count = 10 //默认10行
	if pars["count"] != "" {
		count, _ = strconv.Atoi(pars["count"])
	}

	p := mParam{
		Type:   pars["type"],
		Offset: offset,
		Count:  count,
	}
	d, err := json.Marshal(p)
	if err != nil {
		log.Error("转换素材列表参数失败,", err)
		return nil, err
	}
	res, err := t.request.Do(api, params, bytes.NewBuffer(d))
	if err != nil {
		log.Error("获取素材列表失败：", err.Error())
		return nil, err
	}
	return res["item"], nil
}

// GetMaterialCount 获取素材总和
func (t *WxClient) GetMaterialCount() (map[string]interface{}, error) {
	api := API["material"]["count"]
	params := url.Values{}
	params.Set("access_token", t.authorizerAccessToken)
	res, err := t.request.Do(api, params)
	if err != nil {
		log.Error("获取素材列表失败：", err.Error())
		return nil, err
	}
	return res, nil
}

// AddMaterial 新增永久素材
func (t *WxClient) AddMaterial(pars []byte, mtype string) (map[string]interface{}, error) {
	api := API["material"]["count"]
	params := url.Values{}
	params.Set("access_token", t.authorizerAccessToken)
	params.Set("type", mtype)
	res, err := t.request.Do(api, params, bytes.NewBuffer(pars))
	if err != nil {
		log.Error("新增永久素材失败：", err.Error())
		return nil, err
	}
	return res, nil
}
