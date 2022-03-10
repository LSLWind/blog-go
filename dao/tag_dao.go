package dao

import "blog-go/models"

type TagDao struct {
}

func (t TagDao) QueryAll() ([]models.Tag, int64) {
	var resp []models.Tag
	globalDb.Find(&resp)

	return resp, int64(len(resp))

}
