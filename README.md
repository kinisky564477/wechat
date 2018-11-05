# 微信公众号接口

关于重复轮子的说明:
1. 想趁着这个机会学习下 GO
2. 其实 Github 上微信公众号接口的所有实现，我感觉都不怎么满意
3. 目前功能很少, 后期项目遇到再一点点添加

## 简易使用说明
**安装**
```bash
go get "github.com/zjxpcyc/wechat"
```


```go
import "github.com/zjxpcyc/wechat"

...

// 微信开发者信息
certificate := map[string]string{
	"appid"  "",
	"secret" "",
	"token" "",
	"aeskey": "",
}

// 一个日志记录器, 需要实现如下接口
// type Log interface {
// 	Critical(string, ...interface{})
// 	Error(string, ...interface{})
// 	Warning(string, ...interface{})
// 	Info(string, ...interface{})
// 	Debug(string, ...interface{})
// }
log := &AnLogger{}

// 初始化, 之后就可以正常使用了
wx := wechat.NewClient(certificate, log)
```

**示例**

1. 首次接入
```golang
func SomeControllerMethod() {
  // fromRequest 代表为微信发送的 http get 请求
  signature := fromRequest.Query("signature")
  timestamp := fromRequest.Query("timestamp")
  nonce := fromRequest.Query("nonce")
  echostr := fromRequest.Query("echostr")

  if wx.Signature(timestamp, nonce) == signature {
    response(echostr)
  }
}
```

2. 获取用户 OpenID
```golang
func SomeControllerMethod() {
  // fromRequest 代表为微信发送的 http get 请求
  // code 需要前端传送过来
  code := fromRequest.Query("code")

  openID, err := wx.GetOpenID(code)
  if err != nil {
    // TODO something
  }

  response(openID)
}
```

3. 获取用户详细信息
```golang
func SomeControllerMethod() {
  // fromRequest 代表为微信发送的 http get 请求
  // code 需要前端传送过来
  code := fromRequest.Query("code")

  user, err := wx.GetUserInfo(code)
  if err != nil {
    // TODO something
  }

  response(user)
}
```
