package contract

import "github.com/gifff/xenelectronic/entity"

// CategoryService contract
type CategoryService interface {
	ListAllCategories() ([]entity.Category, error)
}
