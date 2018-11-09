package component

import (
	"github.com/kinisky564477/wechat/core"
	"github.com/kinisky564477/wechat/wx"
)

// WxClient 微信客户端
type WxClient struct {
	authorizerAccessToken   string
	authorizerRefreshToken  string
	authorizationCode       string
	kernel                  *core.Kernel
	request                 core.Request
	getComponentToken       func() string
	getComponentCertificate func() map[string]string

	/*
	* certificate key 值如下:
	* 	appid 	开发者ID(AppID)
	*		secret 	开发者密码(AppSecret)
	*   token 	令牌(Token)
	*   aeskey 	消息加解密密钥 (EncodingAESKey)
	 */
	certificate map[string]string

	/*
		微信授权返回值，授权码及授权列表
	*/
	authorizerResult map[string]interface{}
}

// NewWxClient 初始客户端
func NewWxClient(certificate map[string]string, getComponentToken func() string, getComponentCertificate func() map[string]string) *WxClient {
	cli := &WxClient{
		certificate:             certificate,
		request:                 core.NewDefaultRequest(wx.CheckJSONResult),
		kernel:                  core.NewKernel(),
		authorizerAccessToken:   certificate["authorizer_access_token"],
		authorizerRefreshToken:  certificate["authorizer_refresh_token"],
		authorizationCode:       certificate["authorization_code"],
		getComponentToken:       getComponentToken,
		getComponentCertificate: getComponentCertificate,
	}

	if cli.authorizerRefreshToken == "" {
		err := cli.AuthorizerToken()
		if err != nil {
			log.Error("微信客户端授权失败：", err)
			return nil
		}
	}

	cli.kernel.SetTask("refresh-token", cli.RefreshToken)
	cli.kernel.StartTask("refresh-token")

	return cli
}