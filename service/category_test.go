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
