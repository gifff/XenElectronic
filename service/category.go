package service

import (
	"github.com/gifff/xenelectronic/contract"
	"github.com/gifff/xenelectronic/entity"
)

var _ contract.CategoryService = (*Category)(nil)

func NewCategory(categoryRepo contract.CategoryRepository) *Category {
	return &Category{
		categoryRepo: categoryRepo,
	}
}

type Category struct {
	categoryRepo contract.CategoryRepository
}

func (c *Category) ListAllCategories() ([]entity.Category, error) {
	return c.categoryRepo.ListAll()
}
