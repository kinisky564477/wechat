package component

import (
	"net/http"

	"github.com/kinisky564477/wechat/core"
)

// API 接口列表
var API = map[string]map[string]core.API{
	"component_token": map[string]core.API{
		"post": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/component/api_component_token",
			ResponseType: "json",
		},
	},
	"pre_auth_code": map[string]core.API{
		"post": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=COMPONENT_ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
	"authorizer_access_token": map[string]core.API{
		"post": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=COMPONENT_ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
	"refresh_access_token": map[string]core.API{
		"post": core.API{
			Method:       http.MethodPost,
			URI:          "https:// api.weixin.qq.com /cgi-bin/component/api_authorizer_token?component_access_token=COMPONENT_ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
}
