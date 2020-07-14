package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gifff/xenelectronic/entity"
	"github.com/gifff/xenelectronic/mocks"
	"github.com/gifff/xenelectronic/service"
)

func TestCreateCart(t *testing.T) {
	cartRepo := &mocks.CartRepository{}
	s := service.NewCart(cartRepo)

	uuid := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	cartRepo.On("CreateCart").Return(uuid, nil)
	gotCartID, gotErr := s.CreateCart()
	wantCartID := uuid

	assert.Nil(t, gotErr)
	assert.Equal(t, wantCartID, gotCartID)
}

func TestListProductsInCart(t *testing.T) {
	cartRepo := &mocks.CartRepository{}
	s := service.NewCart(cartRepo)

	uuid := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	cartRepo.On("ListProductsByCartID", uuid).Return([]entity.CartItem{
		{
			ID:     3000,
			CartID: uuid,
			Product: entity.Product{
				ID:          2020,
				CategoryID:  1010,
				Name:        "iPhone 11",
				Description: "Apple Smartphone",
				Price:       13000000,
			},
		},
		{
			ID:     3001,
			CartID: uuid,
			Product: entity.Product{
				ID:          2222,
				CategoryID:  1010,
				Name:        "Samsung Galaxy S20 Ultra",
				Description: "Samsung Smartphone",
				Price:       20000000,
			},
		},
	}, nil)
	gotCartItems, gotErr := s.ListProductsInCart(uuid)
	wantCartItems := []entity.CartItem{
		{
			ID:     3000,
			CartID: uuid,
			Product: entity.Product{
				ID:          2020,
				CategoryID:  1010,
				Name:        "iPhone 11",
				Description: "Apple Smartphone",
				Price:       13000000,
			},
		},
		{
			ID:     3001,
			CartID: uuid,
			Product: entity.Product{
				ID:          2222,
				CategoryID:  1010,
				Name:        "Samsung Galaxy S20 Ultra",
				Description: "Samsung Smartphone",
				Price:       20000000,
			},
		},
	}

	assert.Nil(t, gotErr)
	assert.Equal(t, wantCartItems, gotCartItems)
}

func TestAddProductIntoCart(t *testing.T) {
	cartRepo := &mocks.CartRepository{}
	s := service.NewCart(cartRepo)

	uuid := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	cartRepo.On("AddProductIntoCart", uuid, int64(2001)).Return(entity.CartItem{
		ID:     3456,
		CartID: uuid,
		Product: entity.Product{
			ID:          2001,
			CategoryID:  1010,
			Name:        "iPhone XR",
			Description: "Apple Smartphone",
			Price:       11000000,
		},
	}, nil)
	gotCartItem, gotErr := s.AddProductIntoCart(uuid, 2001)
	wantCartItem := entity.CartItem{
		ID:     3456,
		CartID: uuid,
		Product: entity.Product{
			ID:          2001,
			CategoryID:  1010,
			Name:        "iPhone XR",
			Description: "Apple Smartphone",
			Price:       11000000,
		},
	}

	assert.Nil(t, gotErr)
	assert.Equal(t, wantCartItem, gotCartItem)
}
