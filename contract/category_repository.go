package contract

import "github.com/gifff/xenelectronic/entity"

type CategoryRepository interface {
	ListAll() ([]entity.Category, error)
}
