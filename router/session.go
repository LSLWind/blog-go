package router

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LSL session 管理 -扩充任意session结构体对象
type Lsl struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//获取当前session
func getCurrentLsl(c *gin.Context) (lsl Lsl) {
	session := sessions.Default(c)
	lsl = session.Get("currentUser").(Lsl) // 类型转换一下
	return lsl
}

func setCurrentUser(c *gin.Context, userInfo User) {
	session := sessions.Default(c)
	session.Set("currentUser", userInfo)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	session.Save()
}

func setupRouter(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		var loginVo User
		if c.ShouldBindJSON(&loginVo) != nil {
			c.String(http.StatusOK, "参数错误")
			return
		}
		if loginVo.Email == db.Email && loginVo.Password == db.Password {
			setCurrentUser(c, *db) // 邮箱和密码正确则将当前用户信息写入session中
			c.String(http.StatusOK, "登录成功")
		} else {
			c.String(http.StatusOK, "登录失败")
		}
	})

	r.GET("/sayHello", func(c *gin.Context) {
		userInfo := getCurrentUser(c)
		c.String(http.StatusOK, "Hello "+userInfo.Name)
	})
}

var db = &User{Id: 10001, Email: "abc@gmail.cn", Username: "Alice", Password: "123456"} // 不操作数据库，把所有用户信息写死在代码里

func main() {
	gob.Register(User{}) // 注册User结构体
	r := gin.Default()
	store := cookie.NewStore([]byte("snaosnca"))
	r.Use(sessions.Sessions("SESSIONID", store))
	setupRouter(r)
	r.Run(":8080")
}
