package dao

import "blog-go/models"

type CategoryDao struct {
}

func (c CategoryDao) QueryAll() ([]models.Category, int64) {
	var resp []models.Category
	globalDb.Table("category").Find(&resp)
	return resp, int64(len(resp))
}
