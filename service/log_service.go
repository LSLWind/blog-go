package service

import (
	"blog-go/dao"
	"blog-go/models"
	"blog-go/models/req"
	"blog-go/models/response"
	"time"
)

type LogService struct {
	logDao dao.LogDao
}

func (l LogService) GetAllLogsContent() ([]response.LogResponse, int64) {
	resp, resLen := l.logDao.QueryAll()
	return convert(resp), resLen
}
func (l LogService) InsertOneLog(request req.LogRequest) bool {
	return l.logDao.InsertOne(convertLogRequest2Log(request))
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

/**
请求数据转换，默认日期
*/
func convertLogRequest2Log(request req.LogRequest) models.Log {
	log := models.Log{}
	log.UpdateTime = time.Now()
	log.CreateTime = time.Now()
	log.Content = request.Content
	log.Type = request.Type
	return log
}
