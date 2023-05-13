package service

import "sribuu/products/entity"

type CategoryService interface {
	Save(entity.Category) entity.Category
	FindAll() []entity.Category
}

type categoryService struct {
	categories []entity.Category
}

func New() CategoryService {
	return &categoryService{}
}

func (service *categoryService) Save(category entity.Category) entity.Category {
	service.categories = append(service.categories, category)

	return category
}

func (service *categoryService) FindAll() []entity.Category {
	return service.categories
}
