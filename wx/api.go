package wx

import (
	"net/http"

	"github.com/zjxpcyc/wechat/core"
)

// API 接口列表
var API = map[string]map[string]core.API{
	"access_token": map[string]core.API{
		"get": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET",
			ResponseType: "json",
		},
	},
	"oauth2": map[string]core.API{
		"access_token": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code",
			ResponseType: "json",
		},
		"refresh_token": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN",
			ResponseType: "json",
		},
		"auth": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID",
			ResponseType: "json",
		},
		"userinfo": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN",
			ResponseType: "json",
		},
	},
	"qrcode": map[string]core.API{
		"create": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
	"user": map[string]core.API{
		"detail": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN",
			ResponseType: "json",
		},
	},
	"tpl_message": map[string]core.API{
		"send": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
	"menu": map[string]core.API{
		"create": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
		"delete": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
		"get": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
	"jssdk": map[string]core.API{
		"ticket": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=jsapi",
			ResponseType: "json",
		},
	},
}
