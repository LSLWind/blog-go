package dao

import (
	"blog-go/models"
	"github.com/sirupsen/logrus"
)

type CategoryDao struct {
}

func (c CategoryDao) QueryAll() ([]models.Category, int64) {
	var resp []models.Category
	globalDb.Table("category").Find(&resp)
	return resp, int64(len(resp))
}

// 添加一级/二级目录，返回主键id
func (c CategoryDao) InsertCategory(category models.Category) uint {
	err := globalDb.Table("category").Create(&category).Error
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Dao插入一级/二级目录失败：", category)
		return 0
	}
	return category.Id
}
