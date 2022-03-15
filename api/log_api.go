package api

import (
	"blog-go/models/req"
	"blog-go/resp"
	"blog-go/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type LogApi struct {
	logService service.LogService
}

func (l LogApi) GetAllLogsContent(c *gin.Context) {
	logs, _ := l.logService.GetAllLogsContent()
	resp.OK(c, logs)
}

func (l LogApi) InsertOneLog(c *gin.Context) {
	logRequest := req.LogRequest{}
	err := c.Bind(&logRequest)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("LogApi绑定日志参数表单失败")
	}

	res := l.logService.InsertOneLog(logRequest)
	if res {

		resp.OK(c, "插入日志成功")
	} else {
		resp.Error(c, "插入日志失败")
	}
}
