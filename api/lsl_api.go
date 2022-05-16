package api

import (
	"blog-go/resp"
	"github.com/gin-gonic/gin"
)

type LslApi struct {
}

// lsl blog 后台登录校验
func (l LslApi) LslCheckout(c *gin.Context) {
	// 获取post 表单输入
	name := c.PostForm("name")
	password := c.PostForm("password")
	logger.Info("登录校验，name={ " + name + "}" + "，password={ " + password + " }")
	//校验用户名与密码
	if !(name == "baopowodedoushisb" && password == "baopowodedoushisb") {
		resp.OK(c, "error")
	} else {
		resp.OK(c, "right")
	}
}
