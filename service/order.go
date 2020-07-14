package service

import (
	"github.com/gifff/xenelectronic/contract"
	"github.com/gifff/xenelectronic/entity"
)

var _ contract.OrderService = (*Order)(nil)

func NewOrder(orderRepo contract.OrderRepository) *Order {
	return &Order{
		orderRepo: orderRepo,
	}
}

type Order struct {
	orderRepo contract.OrderRepository
}

func (c *Order) Checkout(cartID, customerName, customerEmail, customerAddress string) (entity.Order, error) {
	order, err := c.orderRepo.CheckoutFromCart(cartID, customerName, customerEmail, customerAddress)
	if err != nil {
		return entity.Order{}, err
	}

	for _, cartItem := range order.CartItems {
		order.PaymentAmount += cartItem.Product.Price
	}
	order.PaymentMethod = "Bank Transfer"
	order.PaymentAccountNumber = "232 555 8965"

	return order, nil
}
