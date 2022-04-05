package dao

import "blog-go/models"

type NavigationDao struct {
}

func (n NavigationDao) QueryAll() ([]models.Navigation, int64) {
	var resp []models.Navigation
	globalDb.Find(&resp)

	return resp, int64(len(resp))

}
