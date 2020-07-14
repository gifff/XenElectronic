package service_test

import (
	"errors"
	"testing"

	"github.com/gifff/xenelectronic/entity"
	"github.com/gifff/xenelectronic/mocks"
	"github.com/gifff/xenelectronic/service"
	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	orderRepo := &mocks.OrderRepository{}
	s := service.NewOrder(orderRepo)

	cartID := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"
	orderID := "deadbeef-dead-beef-dead-beefdeadbeef"

	orderRepo.On("CheckoutFromCart", cartID, "John Doe", "john.doe@example.com", "1 Hacker Way").Return(entity.Order{
		ID:              orderID,
		CustomerName:    "John Doe",
		CustomerEmail:   "john.doe@example.com",
		CustomerAddress: "1 Hacker Way",
		CartItems: []entity.CartItem{
			{
				ID:     3300,
				CartID: cartID,
				Product: entity.Product{
					ID:          2020,
					CategoryID:  1010,
					Name:        "iPhone 11",
					Description: "Apple Smartphone",
					Price:       13000000,
				},
			},
			{
				ID:     3330,
				CartID: cartID,
				Product: entity.Product{
					ID:          2222,
					CategoryID:  1010,
					Name:        "Samsung Galaxy S20 Ultra",
					Description: "Samsung Smartphone",
					Price:       20000000,
				},
			},
			{
				ID:     3333,
				CartID: cartID,
				Product: entity.Product{
					ID:          2000,
					CategoryID:  1010,
					Name:        "iPhone X",
					Description: "Apple Smartphone",
					Price:       10000000,
				},
			},
		},
	}, nil)
	gotOrderID, gotErr := s.Checkout(cartID, "John Doe", "john.doe@example.com", "1 Hacker Way")
	wantOrderID := entity.Order{
		ID:                   orderID,
		CustomerName:         "John Doe",
		CustomerEmail:        "john.doe@example.com",
		CustomerAddress:      "1 Hacker Way",
		PaymentAmount:        43000000,
		PaymentMethod:        "Bank Transfer",
		PaymentAccountNumber: "232 555 8965",
		CartItems: []entity.CartItem{
			{
				ID:     3300,
				CartID: cartID,
				Product: entity.Product{
					ID:          2020,
					CategoryID:  1010,
					Name:        "iPhone 11",
					Description: "Apple Smartphone",
					Price:       13000000,
				},
			},
			{
				ID:     3330,
				CartID: cartID,
				Product: entity.Product{
					ID:          2222,
					CategoryID:  1010,
					Name:        "Samsung Galaxy S20 Ultra",
					Description: "Samsung Smartphone",
					Price:       20000000,
				},
			},
			{
				ID:     3333,
				CartID: cartID,
				Product: entity.Product{
					ID:          2000,
					CategoryID:  1010,
					Name:        "iPhone X",
					Description: "Apple Smartphone",
					Price:       10000000,
				},
			},
		},
	}

	assert.Nil(t, gotErr)
	assert.Equal(t, wantOrderID, gotOrderID)
}

func TestCheckout_Error(t *testing.T) {
	orderRepo := &mocks.OrderRepository{}
	s := service.NewOrder(orderRepo)

	cartID := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	orderRepo.On("CheckoutFromCart", cartID, "John Doe", "john.doe@example.com", "1 Hacker Way").Return(entity.Order{}, errors.New("unable to connect to database"))
	gotOrderID, gotErr := s.Checkout(cartID, "John Doe", "john.doe@example.com", "1 Hacker Way")
	wantOrderID := entity.Order{}

	assert.Equal(t, errors.New("unable to connect to database"), gotErr)
	assert.Equal(t, wantOrderID, gotOrderID)
}
