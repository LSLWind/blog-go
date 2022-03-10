package api

import (
	"blog-go/resp"
	"blog-go/service"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
	logService service.LogService
}

func (l LogApi) GetAllLogsContent(c *gin.Context) {
	logs, _ := l.logService.GetAllLogsContent()
	resp.OK(c, logs)
}
