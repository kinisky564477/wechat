package wx_test

import (
	"testing"

	"github.com/kinisky564477/wechat/wx"
)

func TestJsTicketSignature(t *testing.T) {
	url := "http://mp.weixin.qq.com?params=value"
	noncestr := "Wm3WZYTPz0wzccnW"
	timestamp := "1414587457"
	ticket := "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg"

	expected := "0f9de62fce790f9a083d5c99e95740ceb90c27ed"

	res := wx.JsTicketSignature(url, noncestr, ticket, timestamp)

	if res != expected {
		t.Fatalf("TestJsTicketSignature fail:  %s - %s", res, expected)
	}
}
