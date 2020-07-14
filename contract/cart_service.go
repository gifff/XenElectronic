package contract

import "github.com/gifff/xenelectronic/entity"

// CartService contract
type CartService interface {
	CreateCart() (string, error)
	ListProductsInCart(cartID string) ([]entity.CartItem, error)
	AddProductIntoCart(cartID string, productID int64) (entity.CartItem, error)
	RemoveProductFromCart(cartID string, productID int64) error
}
