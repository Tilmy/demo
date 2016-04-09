package Login

import (
	. "github.com/lessgo/lessgo"
)

var IndexHandle = DescHandler{
	Desc: "后台管理登录操作",
	Param: map[string]string{
		"user":     "string | 用户名",
		"password": "string | 密码",
	},
	Handler: func(ctx Context) error {
		return ctx.Render(200,
			"SystemView/Admin/Login/index.tpl",
			map[string]interface{}{
				"name":       ctx.Param("user"),
				"password":   ctx.Param("password"),
				"repeatfunc": repeatfunc,
			})
	},
}
