package wx

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// TplMessageData 模板消息节点
type TplMessageData struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// MiniProgramData 小程序内容
type MiniProgramData struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath,omitempty"`
}

// TplMessage 模板消息
type TplMessage struct {
	ToUser      string                    `json:"touser"`
	TemplateID  string                    `json:"template_id"`
	URL         string                    `json:"url"`
	MiniProgram *MiniProgramData          `json:"miniprogram,omitempty"`
	Data        map[string]TplMessageData `json:"data"`
}

// SendTplMessage 发送模板消息
func (t *Client) SendTplMessage(to, tplID, link string, data map[string]TplMessageData) error {
	api := API["tpl_message"]["send"]
	params := url.Values{}
	params.Set("access_token", t.accessToken)

	message := TplMessage{
		ToUser:     to,
		TemplateID: tplID,
		URL:        link,
		Data:       data,
	}

	b, err := json.Marshal(message)
	if err != nil {
		log.Error("转换模板消息失败", err.Error())
		return err
	}

	_, err = t.request.Do(api, params, bytes.NewBuffer(b))
	if err != nil {
		log.Error("发送模板消息失败", err.Error())
		return err
	}

	return nil
}
