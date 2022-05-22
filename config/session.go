package config

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Session struct {
}

// LSL session 管理 -扩充任意session结构体对象
type Lsl struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//获取当前session
func (s Session) GetLslSession(c *gin.Context) (lsl Lsl) {
	session := sessions.Default(c)
	lsl = session.Get("currentUser").(Lsl) // 类型转换一下
	return lsl
}

// 设置指定字段session
func (s Session) SetLslSession(c *gin.Context, lsl Lsl) {
	session := sessions.Default(c)
	session.Set("currentUser", lsl)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	err := session.Save()
	if err != nil {
		logger.Error(err)
	}

}

// 清空所有session
func (s Session) Logout(c *gin.Context) {
	//clear session
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		logger.Error(err)
	}
	return
}
