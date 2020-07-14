package contract

import "github.com/gifff/xenelectronic/entity"

// OrderRepository contract
type OrderRepository interface {
	CheckoutFromCart(cartID, customerName, customerEmail, customerAddress string) (entity.Order, error)
}
