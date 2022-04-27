package service

import (
	"blog-go/service"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllCategory(t *testing.T) {
	c := service.CategoryService{}
	fmt.Println(c.GetAllCategory())
	assert.NotEmpty(t, c.GetAllCategory())
}
