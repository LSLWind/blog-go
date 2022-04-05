package api

import (
	"blog-go/resp"
	"blog-go/service"
	"github.com/gin-gonic/gin"
)

type NavigationApi struct {
	navigationService service.NavigationService
}

func (n NavigationApi) GetAllNavigationInfo(c *gin.Context) {
	infoResponses := n.navigationService.GetAllNavigationInfo()
	resp.OK(c, infoResponses)
}
