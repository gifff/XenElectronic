package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
