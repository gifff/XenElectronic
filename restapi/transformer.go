package restapi

import (
	"github.com/gifff/xenelectronic/entity"
	"github.com/gifff/xenelectronic/models"
	"github.com/go-openapi/strfmt"
)

func transformOrder(order entity.Order) *models.Order {
	custEmail := strfmt.Email(order.CustomerEmail)
	orderModel := &models.Order{
		ID:                   strfmt.UUID(order.ID),
		CustomerName:         &order.CustomerName,
		CustomerEmail:        &custEmail,
		CustomerAddress:      &order.CustomerAddress,
		PaymentAmount:        order.PaymentAmount,
		PaymentMethod:        order.PaymentMethod,
		PaymentAccountNumber: order.PaymentAccountNumber,
		CartItems:            make(models.CartItems, len(order.CartItems)),
	}
	for i := range order.CartItems {
		p := &models.Product{
			ID:          order.CartItems[i].Product.ID,
			CategoryID:  &order.CartItems[i].Product.CategoryID,
			Name:        &order.CartItems[i].Product.Name,
			Description: &order.CartItems[i].Product.Description,
			Photo:       order.CartItems[i].Product.Photo,
			Price:       &order.CartItems[i].Product.Price,
		}

		orderModel.CartItems[i] = &models.CartItem{
			ID:        order.CartItems[i].ID,
			Product:   p,
			ProductID: &order.CartItems[i].Product.ID,
		}
	}

	return orderModel
}
