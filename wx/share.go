package wx

import (
	"net/url"
	"strconv"
	"time"

	"github.com/kinisky564477/wechat/core"
)

// JsTicketTask 刷新 任务
func (t *Client) JsTicketTask() time.Duration {
	var reTrySec int64 = 60
	ticket, expire, err := t.getJsTicket()
	if err != nil {
		log.Error("获取 JS Ticket 失败", err.Error())
		expire = reTrySec
	}

	// 不允许连续不断调用
	if expire == 0 {
		expire = reTrySec
	}

	t.jsTicket = ticket
	return time.Duration(expire) * time.Second
}

// getJsTicket 获取 token
func (t *Client) getJsTicket() (string, int64, error) {
	api := API["jssdk"]["ticket"]
	params := url.Values{}
	params.Set("access_token", t.accessToken)

	res, err := t.request.Do(api, params)
	if err != nil {
		return "", 0, err
	}

	ticket, _ := res["ticket"].(string)
	expire, _ := res["expires_in"].(float64)
	return ticket, int64(expire), nil
}

// GetJsTicketSignature js-sdk signature
func (t *Client) GetJsTicketSignature(url string) map[string]string {
	noncestr := core.RandLimitString(16)
	timestamp := strconv.FormatInt(time.Now().Local().Unix(), 10)

	signature := JsTicketSignature(url, noncestr, t.jsTicket, timestamp)

	return map[string]string{
		"noncestr":  noncestr,
		"timestamp": timestamp,
		"signature": signature,
		"appId":     t.certificate["appid"],
	}
}
