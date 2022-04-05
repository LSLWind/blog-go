package service

import (
	"blog-go/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllNavigationInfo(t *testing.T) {
	n := service.NavigationService{}
	assert.NotNil(t, n.GetAllNavigationInfo())
}
