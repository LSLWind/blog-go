package dao

import "blog-go/models"

type LogDao struct {
}

func (l LogDao) QueryAll() ([]models.Log, int64) {
	var resp []models.Log
	globalDb.Find(&resp)
	return resp, int64(len(resp))
}
