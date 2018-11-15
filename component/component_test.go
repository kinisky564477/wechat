package component_test

import (
	"testing"

	"github.com/astaxie/beego"

	"github.com/kinisky564477/wechat/component"
)

func TestNewComponentClient(t *testing.T) {
	var cert = map[string]string{
		"appid":  "wx9fd33312e78e8d02",
		"aeskey": "41b2994de43d4e3b9dc0f54ee8c5c1bb0050496367e",
		"secret": "69ee34668cd2635138b831f9ecb1fb4f",
	}
	cli := component.NewComponentClient(cert)
	ti := cli.ComponentAccessTokenTask()
	beego.Error(ti)
}
