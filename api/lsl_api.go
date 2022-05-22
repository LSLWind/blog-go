package api

import (
	"blog-go/config"
	"blog-go/models/req"
	"blog-go/resp"
	"blog-go/service"
	"github.com/gin-gonic/gin"
)

type LslApi struct {
	session         config.Session
	categoryService service.CategoryService
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
		// 保存session
		lsl := config.Lsl{}
		lsl.Name = "baopowodedoushisb"
		lsl.Password = "baopowodedoushisb"
		l.session.SetLslSession(c, lsl)
		resp.OK(c, "right")
	}
}

// 增肌一级目录
func (l LslApi) AddOneLevelCategory(c *gin.Context) {
	if !l.checkLslSession(c) {
		resp.Error(c, "session校验失败")
	}
	//校验参数
	title := c.PostForm("title")
	subTitle := c.PostForm("subTitle")
	svgIcon := c.PostForm("svgIcon")
	if subTitle == "" {
		resp.Error(c, "二级目录不能为空")
	}
	//插入
	oneLevelCategoryRequest := req.OneLevelCategoryRequest{
		Title:    title,
		SubTitle: subTitle,
		SvgIcon:  svgIcon,
	}

	l.categoryService.AddOneLevelCategory(oneLevelCategoryRequest)

	resp.OK(c)
}

// 校验Lsl session
func (l LslApi) checkLslSession(c *gin.Context) bool {
	lsl := l.session.GetLslSession(c)
	logger.Info("校验session" + lsl.Password)
	if lsl.Name != "baopowodedoushisb" || lsl.Password != "baopowodedoushisb" {
		return false
	}
	return true
}
