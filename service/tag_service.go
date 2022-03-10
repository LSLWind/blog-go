package service

import (
	"blog-go/dao"
	"blog-go/models"
)

type TagService struct {
	tagDao dao.TagDao
}

func (s TagService) QueryAll() ([]models.Tag, int64) {
	return s.tagDao.QueryAll()
}
