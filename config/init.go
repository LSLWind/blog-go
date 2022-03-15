package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New() // 创建一个log示例
func init() {
	//日志框架logrus初始化
	err := initLogrus()
	if err != nil {
		fmt.Println(err)
	}
	logger.Info("初始化日志框架logrus成功")

}

func initLogrus() error {
	logger.Formatter = &logrus.JSONFormatter{}                                           // 设置为json格式的日志
	file, err := os.OpenFile("./gin_log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // 创建一个log日志文件
	if err != nil {
		fmt.Println("创建文件/打开文件失败！")
		return err
	}
	logger.Out = file               // 设置log的默认文件输出
	gin.SetMode(gin.ReleaseMode)    // 发布版本
	gin.DefaultWriter = logger.Out  // gin框架自己记录的日志也会输出
	logger.Level = logrus.InfoLevel // 设置日志级别
	return nil
}
