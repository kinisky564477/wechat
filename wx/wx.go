package wx

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"time"

	"github.com/zjxpcyc/wechat/core"
)

var log core.Log

// Client 微信公众号接口客户端
type Client struct {
	accessToken string
	jsTicket    string
	kernel      *core.Kernel
	request     core.Request

	/*
	* certificate key 值如下:
	* 	appid 	开发者ID(AppID)
	*		secret 	开发者密码(AppSecret)
	*   token 	令牌(Token)
	*   aeskey 	消息加解密密钥 (EncodingAESKey)
	 */
	certificate map[string]string
}

// NewClient 初始客户端
func NewClient(certificate map[string]string) *Client {
	cli := &Client{
		request:     core.NewDefaultRequest(checkJSONResult),
		kernel:      core.NewKernel(),
		certificate: certificate,
	}

	cli.kernel.SetTask("access-token", cli.AccessTokenTask)
	cli.kernel.StartTask("access-token")

	cli.kernel.SetTask("js-ticket", cli.JsTicketTask)
	cli.kernel.StartTask("js-ticket", 30*time.Second)

	return cli
}

// GetAppID 获取 AppID
func (t *Client) GetAppID() string {
	return t.certificate["appid"]
}

// Signature 初始校验
func (t *Client) Signature(timestamp, nonce string) string {
	token := t.certificate["token"]
	strs := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SetLogInst 设置全局日志实例
func SetLogInst(l core.Log) {
	core.SetLogInst(l)
	log = l
}

func init() {
	log = &core.DefaultLogger{}
}
