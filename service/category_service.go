package service

import (
	"blog-go/dao"
	"blog-go/models"
	"blog-go/models/req"
	"blog-go/models/response"
	"time"
)

type CategoryService struct {
	categoryDao dao.CategoryDao
}

func (c CategoryService) GetAllCategory() []response.CategoryResponse {
	resp, _ := c.categoryDao.QueryAll()
	return category2CategoryResponse(resp)
}

// 增加一级目录xinxi
func (c CategoryService) AddOneLevelCategory(category req.OneLevelCategoryRequest) bool {
	// 先插入二级目录
	twoLevelCategory := models.Category{
		Name:       category.SubTitle,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	subId := c.categoryDao.InsertCategory(twoLevelCategory)
	//插入一级目录
	oneLevelCategory := models.Category{
		Name:            category.Title,
		SubCategoryId:   subId,
		SubCategoryName: category.SubTitle,
		SvgIcon:         category.SvgIcon,
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}
	id := c.categoryDao.InsertCategory(oneLevelCategory)

	if id == 0 {
		logger.Error("插入一级目录失败")
		return false
	}
	return true
}

/**
返回数据转换，二级目录格式
*/
func category2CategoryResponse(categories []models.Category) []response.CategoryResponse {
	categoryResponses := []response.CategoryResponse{}

	for _, category := range categories {
		if category.SubCategoryId != 0 {
			isNew := true
			//遍历一级目录，如果是一级目录则在该目录下加入二级目录
			for i, res := range categoryResponses {
				if category.Name == res.Name {
					isNew = false
					subCategory := response.SubCategory{}
					subCategory.Id = category.SubCategoryId
					subCategory.Name = category.SubCategoryName
					res.SubCategoryList = append(res.SubCategoryList, subCategory)
					categoryResponses[i] = res
				}
			}
			//增加一条新的一级目录
			if isNew {
				categoryResponse := response.CategoryResponse{}
				categoryResponse.Id = category.Id
				categoryResponse.Name = category.Name
				categoryResponse.SvgIcon = category.SvgIcon
				//一级目录下的第一条子目录
				subCategory := response.SubCategory{}
				subCategory.Id = category.SubCategoryId
				subCategory.Name = category.SubCategoryName
				categoryResponse.SubCategoryList = append(categoryResponse.SubCategoryList, subCategory)
				categoryResponses = append(categoryResponses, categoryResponse)
			}
		}
	}
	return categoryResponses
}

//func oneLevelCategoryReq2Category(request req.OneLevelCategoryRequest)models.Category  {
//	category := models.Category{}
//	category.Name = request.Title
//	category.SubCategoryName =
//}
