package contract

import "github.com/gifff/xenelectronic/entity"

// OrderService contract
type OrderService interface {
	Checkout(cartID, customerName, customerEmail, customerAddress string) (entity.Order, error)
	View(orderID string) (entity.Order, error)
}
