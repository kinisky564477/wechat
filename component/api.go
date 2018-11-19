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
			URI:          "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=COMPONENT_ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
	"wechat": map[string]core.API{
		"getinfo": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=COMPONENT_ACCESS_TOKEN",
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
	// 获取素材列表
	"material": map[string]core.API{
		"count": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
		"list": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
		"add": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=TYPE",
			ResponseType: "json",
		},
		"del": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
	// 客服消息
	"custom": map[string]core.API{
		"message": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN",
			ResponseType: "json",
		},
	},
}
