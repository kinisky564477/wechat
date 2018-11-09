package component

import (
	"github.com/kinisky564477/wechat/core"
	"github.com/kinisky564477/wechat/wx"
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
		certificate: certificate,
		request:     core.NewDefaultRequest(wx.CheckJSONResult),
		kernel:      core.NewKernel(),
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
