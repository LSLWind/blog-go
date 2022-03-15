package dao

import (
	"blog-go/models"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

type LogDao struct {
}

func (l LogDao) QueryAll() ([]models.Log, int64) {
	var resp []models.Log
	globalDb.Find(&resp)
	return resp, int64(len(resp))
}

func (l LogDao) InsertOne(log models.Log) bool {
	//一般项目中我们会类似下面的写法，通过Error对象检测，插入数据有没有成功，如果没有错误那就是数据写入成功了。
	err := globalDb.Create(&log).Error
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Dao插入日志Log失败")
		return false
	}
	return true
}
