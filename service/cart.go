package service

import (
	"github.com/gifff/xenelectronic/contract"
	"github.com/gifff/xenelectronic/entity"
)

var _ contract.CartService = (*Cart)(nil)

func NewCart(cartRepo contract.CartRepository) *Cart {
	return &Cart{
		cartRepo: cartRepo,
	}
}

type Cart struct {
	cartRepo contract.CartRepository
}

func (c *Cart) CreateCart() (string, error) {
	return c.cartRepo.CreateCart()
}

func (c *Cart) ListProductsInCart(cartID string) ([]entity.CartItem, error) {
	return c.cartRepo.ListProductsByCartID(cartID)
}

func (c *Cart) AddProductIntoCart(cartID string, productID int64) (entity.CartItem, error) {
	return c.cartRepo.AddProductIntoCart(cartID, productID)
}

func (c *Cart) RemoveProductFromCart(cartID string, productID int64) error {
	return c.cartRepo.RemoveProductFromCart(cartID, productID)
}
