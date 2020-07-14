package contract

import "github.com/gifff/xenelectronic/entity"

// CategoryService contract
type CategoryService interface {
	ListAllCategories() ([]entity.Category, error)
	ListProductsByCategoryID(categoryID, since int64, limit int32) ([]entity.Product, error)
}
