package wx

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"
)

type QRCodeAction struct {
	SceneID  int64  `json:"scene_id,omitempty"`
	SceneStr string `json:"scene_str,omitempty"`
}

type QRCodeMessage struct {
	ExpireSeconds int64                   `json:"expire_seconds,omitempty"`
	ActionName    string                  `json:"action_name"`
	ActionInfo    map[string]QRCodeAction `json:"action_info"`
}

type QRCodeResult struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
}

const (
	QRScene         = "QR_SCENE"
	QRStrScene      = "QR_STR_SCENE"
	QRLimitScene    = "QR_LIMIT_SCENE"
	QRLimitStrScene = "QR_LIMIT_STR_SCENE"
)

// GetQRCode 获取二维码
func (t *Client) GetQRCode(message *QRCodeMessage) (*QRCodeResult, error) {
	return t.createQrCode(*message)
}

// GetTempStrQRCode 获取临时字符串参数二维码
func (t *Client) GetTempStrQRCode(message string, expire ...time.Duration) (string, error) {
	// default expire_seconds is a week
	expTime := 7 * 24 * time.Hour
	if len(expire) > 0 {
		expTime = expire[0]
	}

	qrcode := QRCodeMessage{
		ExpireSeconds: int64(expTime),
		ActionName:    QRStrScene,
		ActionInfo: map[string]QRCodeAction{
			"scene": QRCodeAction{
				SceneStr: message,
			},
		},
	}

	res, err := t.createQrCode(qrcode)
	if err != nil {
		log.Error("创建二维码失败", err.Error())
		return "", err
	}

	return res.URL, nil
}

// GetStrQRCode 获取永久字符串参数二维码
func (t *Client) GetStrQRCode(message string) (string, error) {
	qrcode := QRCodeMessage{
		ActionName: QRLimitStrScene,
		ActionInfo: map[string]QRCodeAction{
			"scene": QRCodeAction{
				SceneStr: message,
			},
		},
	}

	res, err := t.createQrCode(qrcode)
	return res.URL, err
}

// GetTempIntQRCode 获取临时字符串参数二维码
func (t *Client) GetTempIntQRCode(message int64, expire ...time.Duration) (string, error) {
	// default expire_seconds is a week
	expTime := 7 * 24 * time.Hour
	if len(expire) > 0 {
		expTime = expire[0]
	}

	qrcode := QRCodeMessage{
		ExpireSeconds: int64(expTime),
		ActionName:    QRScene,
		ActionInfo: map[string]QRCodeAction{
			"scene": QRCodeAction{
				SceneID: message,
			},
		},
	}

	res, err := t.createQrCode(qrcode)
	if err != nil {
		log.Error("创建二维码失败", err.Error())
		return "", err
	}

	return res.URL, err
}

// GetIntQRCode 获取永久字符串参数二维码
func (t *Client) GetIntQRCode(message int64) (string, error) {
	qrcode := QRCodeMessage{
		ActionName: QRLimitScene,
		ActionInfo: map[string]QRCodeAction{
			"scene": QRCodeAction{
				SceneID: message,
			},
		},
	}

	res, err := t.createQrCode(qrcode)
	return res.URL, err
}

// createQrCode 获取字符串参数的二维码
func (t *Client) createQrCode(message QRCodeMessage) (*QRCodeResult, error) {
	api := API["qrcode"]["create"]
	params := url.Values{}
	params.Set("access_token", t.accessToken)

	data, err := json.Marshal(message)
	if err != nil {
		log.Error("转换二维码消息失败,", err)
		return nil, err
	}

	res, err := t.request.Do(api, params, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	expire, _ := res["expire_seconds"].(float64)
	qrRes := QRCodeResult{
		Ticket:        res["ticket"].(string),
		ExpireSeconds: int64(expire),
		URL:           res["url"].(string),
	}

	return &qrRes, nil
}
