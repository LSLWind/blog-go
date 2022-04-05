package service

import (
	"blog-go/dao"
	"blog-go/models"
	"blog-go/models/response"
)

type NavigationService struct {
	navigationDao dao.NavigationDao
}

func (n NavigationService) GetAllNavigationInfo() map[string][]response.NavigationInfoResponse {
	resp, _ := n.navigationDao.QueryAll()
	return navigations2InfoResponses(resp)
}

//按照type进行分类-map
func navigations2InfoResponses(navigations []models.Navigation) map[string][]response.NavigationInfoResponse {
	infoMap := make(map[string][]response.NavigationInfoResponse)

	for _, n := range navigations {
		infoResponse := navigation2InfoResponse(n)
		//将元素加入map中
		if value, exists := infoMap[infoResponse.Type]; exists {
			_ = append(value, infoResponse)
		} else {
			types := []response.NavigationInfoResponse{infoResponse}
			infoMap[infoResponse.Type] = types
		}
	}

	return infoMap
}

func navigation2InfoResponse(navigation models.Navigation) response.NavigationInfoResponse {
	infoResponse := response.NavigationInfoResponse{}
	infoResponse.Name = navigation.Name
	infoResponse.Type = navigation.Type
	infoResponse.LinkUrl = navigation.LinkUrl
	infoResponse.Id = navigation.Id
	return infoResponse
}
