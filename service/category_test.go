package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gifff/xenelectronic/entity"
	"github.com/gifff/xenelectronic/mocks"
	"github.com/gifff/xenelectronic/service"
)

func TestListAllCategories(t *testing.T) {
	categoryRepo := &mocks.CategoryRepository{}
	s := service.NewCategory(categoryRepo)

	categoryRepo.On("ListAll").Return(
		[]entity.Category{
			{
				ID:   1000,
				Name: "Home Appliances",
			},
			{
				ID:   1001,
				Name: "Handphones",
			},
			{
				ID:   1010,
				Name: "Smartphones",
			},
		},
		nil,
	)

	gotCategories, gotErr := s.ListAllCategories()
	wantCategories := []entity.Category{
		{
			ID:   1000,
			Name: "Home Appliances",
		},
		{
			ID:   1001,
			Name: "Handphones",
		},
		{
			ID:   1010,
			Name: "Smartphones",
		},
	}

	assert.Nil(t, gotErr)
	assert.Equal(t, wantCategories, gotCategories)
}

func TestListProductsByCategoryID(t *testing.T) {
	categoryRepo := &mocks.CategoryRepository{}
	s := service.NewCategory(categoryRepo)

	categoryRepo.On("ListProductsByCategoryID", int64(1010), int64(10), int32(5)).Return(
		[]entity.Product{
			{
				ID:          2000,
				CategoryID:  1010,
				Name:        "iPhone X",
				Description: "Apple Smartphone",
				Price:       10000000,
			},
			{
				ID:          2001,
				CategoryID:  1010,
				Name:        "iPhone XR",
				Description: "Apple Smartphone",
				Price:       11000000,
			},
			{
				ID:          2010,
				CategoryID:  1010,
				Name:        "iPhone XS",
				Description: "Apple Smartphone",
				Price:       15000000,
			},
			{
				ID:          2020,
				CategoryID:  1010,
				Name:        "iPhone 11",
				Description: "Apple Smartphone",
				Price:       13000000,
			},
			{
				ID:          2222,
				CategoryID:  1010,
				Name:        "Samsung Galaxy S20 Ultra",
				Description: "Samsung Smartphone",
				Price:       20000000,
			},
		},
		nil,
	)

	gotCategories, gotErr := s.ListProductsByCategoryID(1010, 10, 5)
	wantCategories := []entity.Product{
		{
			ID:          2000,
			CategoryID:  1010,
			Name:        "iPhone X",
			Description: "Apple Smartphone",
			Price:       10000000,
		},
		{
			ID:          2001,
			CategoryID:  1010,
			Name:        "iPhone XR",
			Description: "Apple Smartphone",
			Price:       11000000,
		},
		{
			ID:          2010,
			CategoryID:  1010,
			Name:        "iPhone XS",
			Description: "Apple Smartphone",
			Price:       15000000,
		},
		{
			ID:          2020,
			CategoryID:  1010,
			Name:        "iPhone 11",
			Description: "Apple Smartphone",
			Price:       13000000,
		},
		{
			ID:          2222,
			CategoryID:  1010,
			Name:        "Samsung Galaxy S20 Ultra",
			Description: "Samsung Smartphone",
			Price:       20000000,
		},
	}

	assert.Nil(t, gotErr)
	assert.Equal(t, wantCategories, gotCategories)
}
