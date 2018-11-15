package component

import (
	"errors"

	"github.com/kinisky564477/wechat/core"
)

var log core.Log

// ComponentClient 微信第三方客户端
type ComponentClient struct {
	componentVerifyTicket string
	componentAccessToken  string
	request               core.Request
	kernel                *core.Kernel

	/*
	* certificate key 值如下:
	* 	component_appid
	*		aeskey 	加密key
	 */
	certificate map[string]string
	wxClients   map[string]*WxClient
}

// NewComponentClient 初始客户端
func NewComponentClient(certificate map[string]string) *ComponentClient {
	cli := &ComponentClient{
		certificate:           certificate,
		request:               core.NewDefaultRequest(CheckJSONResult),
		kernel:                core.NewKernel(),
		componentVerifyTicket: certificate["componentVerifyTicket"],
	}
	cli.kernel.SetTask("component-token", cli.ComponentAccessTokenTask)
	cli.kernel.StartTask("component-token")
	return cli
}

// RefreshTicket 刷新ticket
func (t *ComponentClient) RefreshTicket(ticket string) {
	t.componentVerifyTicket = ticket
}

// AppendWxClient 添加wxclient
func (t *ComponentClient) AppendWxClient(wxClient *WxClient) {
	if t.wxClients == nil {
		t.wxClients = map[string]*WxClient{
			wxClient.certificate["appid"]: wxClient,
		}
	} else {
		t.wxClients[wxClient.certificate["appid"]] = wxClient
	}
}

// GetToken 获取token
func (t *ComponentClient) GetToken() string {
	return t.componentAccessToken
}

// GetCertificate 获取参数
func (t *ComponentClient) GetCertificate() map[string]string {
	return t.certificate
}

// GetWxClient 获取微信client
func (t *ComponentClient) GetWxClient(appid string) (*WxClient, error) {
	if len(t.wxClients) > 0 && t.wxClients[appid] != nil {
		return t.wxClients[appid], nil
	}
	return nil, errors.New("没有对应的微信信息！")
}

// SetLogInst 设置全局日志实例
func SetLogInst(l core.Log) {
	core.SetLogInst(l)
	log = l
}

func init() {
	log = &core.DefaultLogger{}
}
