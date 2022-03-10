package service

import (
	"blog-go/dao"
	"blog-go/models"
	"blog-go/models/response"
)

type LogService struct {
	logDao dao.LogDao
}

func (l LogService) GetAllLogsContent() ([]response.LogResponse, int64) {
	resp, resLen := l.logDao.QueryAll()
	return convert(resp), resLen
}

/**
返回数据转换
*/
func convert(logs []models.Log) []response.LogResponse {
	logResponses := make([]response.LogResponse, len(logs))
	for i, l := range logs {
		var logResponse response.LogResponse
		logResponse.Content = l.Content
		logResponse.UpdateTime = l.UpdateTime
		logResponses[i] = logResponse
	}
	return logResponses
}
