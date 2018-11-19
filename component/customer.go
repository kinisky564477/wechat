// 客户接口
package component

import (
	"bytes"
	"encoding/json"
	"net/url"
)

type CustomTextContent struct {
	Content string `json:"content"`
}

// CustomText 文本
type CustomText struct {
	ToUser  string            `json:"touser"`
	MsgType string            `json:"msgtype"`
	Text    CustomTextContent `json:"text"`
}

// SendCustomText 发送客服文本消息
func (t *WxClient) SendCustomText(to, text string) error {
	data := CustomText{
		ToUser:  to,
		MsgType: "text",
		Text:    CustomTextContent{text},
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Error("转换客服文本失败", err)
		return err
	}

	api := API["custom"]["message"]
	params := url.Values{}
	params.Set("access_token", t.GetToken())

	_, err = t.request.Do(api, params, bytes.NewBuffer(b))
	if err != nil {
		log.Error("发送客服文本消息出错", err)
		return err
	}

	return nil
}
