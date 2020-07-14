package contract

import "github.com/gifff/xenelectronic/entity"

type CategoryRepository interface {
	ListAll() ([]entity.Category, error)
	ListProductsByCategoryID(categoryID, since int64, limit int32) ([]entity.Product, error)
}
