package controller

import (
	"sribuu/products/entity"
	"sribuu/products/service"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	FindAll() []entity.Category
	Save(ctx *gin.Context) (entity.Category, error)
}

type controller struct {
	service service.CategoryService
}

func New(service service.CategoryService) CategoryController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Category {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) (entity.Category, error) {
	var category entity.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		return category, err
	}
	return c.service.Save(category), nil

}
