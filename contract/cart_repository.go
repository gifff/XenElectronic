package contract

import "github.com/gifff/xenelectronic/entity"

type CartRepository interface {
	CreateCart() (string, error)
	ListProductsByCartID(cartID string) ([]entity.CartItem, error)
}
