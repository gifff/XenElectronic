package postgresql

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gifff/xenelectronic/entity"
	"github.com/stretchr/testify/assert"
)

func TestCheckoutFromCart(t *testing.T) {
	cartID := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"
	orderID := "deadbeef-dead-beef-dead-beefdeadbeef"

	testCases := []struct {
		name             string
		generatedOrderID string
		cartID           string
		customerName     string
		customerEmail    string
		customerAddress  string
		dbMockFn         func(mock sqlmock.Sqlmock)
		wantOrder        entity.Order
		wantErr          error
	}{
		{
			name:             "success",
			generatedOrderID: orderID,
			cartID:           cartID,
			customerName:     "John Doe",
			customerEmail:    "john.doe@example.com",
			customerAddress:  "1 Hacker Way",
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO orders (id, customer_name, customer_email, customer_address) VALUES ($1, $2, $3, $4) RETURNING id").
					WithArgs(orderID, "John Doe", "john.doe@example.com", "1 Hacker Way").
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).
							AddRow(orderID),
					)
				mock.ExpectQuery(`INSERT INTO order_items (id, order_id, product_id, name, description, photo, price)
				SELECT ci.id, $1, ci.product_id, p.name, p.description, p.photo, p.price
				FROM cart_items ci
				JOIN products p ON p.id = ci.product_id
				WHERE ci.cart_id = $2
				RETURNING id, order_id, product_id, name, description, photo, price`).
					WithArgs(orderID, cartID).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "order_id", "product_id", "name", "description", "photo", "price"}).
							AddRow("3300", orderID, "2020", "iPhone 11", "Apple Smartphone", "", "13000000").
							AddRow("3330", orderID, "2222", "Samsung Galaxy S20 Ultra", "Samsung Smartphone", "", "20000000").
							AddRow("3333", orderID, "2000", "iPhone X", "Apple Smartphone", "", "10000000"),
					)
				mock.ExpectExec("DELETE FROM cart_items WHERE cart_id = $1").
					WithArgs(cartID).
					WillReturnResult(sqlmock.NewResult(0, 3))
				mock.ExpectExec("DELETE FROM carts WHERE cart_id = $1").
					WithArgs(cartID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			wantOrder: entity.Order{
				ID:              orderID,
				CustomerName:    "John Doe",
				CustomerEmail:   "john.doe@example.com",
				CustomerAddress: "1 Hacker Way",
				CartItems: []entity.CartItem{
					{
						ID:     3300,
						CartID: orderID,
						Product: entity.Product{
							ID:          2020,
							Name:        "iPhone 11",
							Description: "Apple Smartphone",
							Price:       13000000,
						},
					},
					{
						ID:     3330,
						CartID: orderID,
						Product: entity.Product{
							ID:          2222,
							Name:        "Samsung Galaxy S20 Ultra",
							Description: "Samsung Smartphone",
							Price:       20000000,
						},
					},
					{
						ID:     3333,
						CartID: orderID,
						Product: entity.Product{
							ID:          2000,
							Name:        "iPhone X",
							Description: "Apple Smartphone",
							Price:       10000000,
						},
					},
				},
			},
		},
		{
			name:             "error when insert into orders",
			generatedOrderID: orderID,
			cartID:           cartID,
			customerName:     "John Doe",
			customerEmail:    "john.doe@example.com",
			customerAddress:  "1 Hacker Way",
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO orders (id, customer_name, customer_email, customer_address) VALUES ($1, $2, $3, $4) RETURNING id").
					WithArgs(orderID, "John Doe", "john.doe@example.com", "1 Hacker Way").
					WillReturnError(sql.ErrTxDone)
			},
			wantOrder: entity.Order{},
			wantErr:   sql.ErrTxDone,
		},
		{
			name:             "error when insert from cart_items into order_items",
			generatedOrderID: orderID,
			cartID:           cartID,
			customerName:     "John Doe",
			customerEmail:    "john.doe@example.com",
			customerAddress:  "1 Hacker Way",
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO orders (id, customer_name, customer_email, customer_address) VALUES ($1, $2, $3, $4) RETURNING id").
					WithArgs(orderID, "John Doe", "john.doe@example.com", "1 Hacker Way").
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).
							AddRow(orderID),
					)
				mock.ExpectQuery(`INSERT INTO order_items (id, order_id, product_id, name, description, photo, price)
				SELECT ci.id, $1, ci.product_id, p.name, p.description, p.photo, p.price
				FROM cart_items ci
				JOIN products p ON p.id = ci.product_id
				WHERE ci.cart_id = $2
				RETURNING id, order_id, product_id, name, description, photo, price`).
					WithArgs(orderID, cartID).
					WillReturnError(sql.ErrTxDone)
			},
			wantOrder: entity.Order{},
			wantErr:   sql.ErrTxDone,
		},
		{
			name:             "error when delete from cart_items",
			generatedOrderID: orderID,
			cartID:           cartID,
			customerName:     "John Doe",
			customerEmail:    "john.doe@example.com",
			customerAddress:  "1 Hacker Way",
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO orders (id, customer_name, customer_email, customer_address) VALUES ($1, $2, $3, $4) RETURNING id").
					WithArgs(orderID, "John Doe", "john.doe@example.com", "1 Hacker Way").
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).
							AddRow(orderID),
					)
				mock.ExpectQuery(`INSERT INTO order_items (id, order_id, product_id, name, description, photo, price)
				SELECT ci.id, $1, ci.product_id, p.name, p.description, p.photo, p.price
				FROM cart_items ci
				JOIN products p ON p.id = ci.product_id
				WHERE ci.cart_id = $2
				RETURNING id, order_id, product_id, name, description, photo, price`).
					WithArgs(orderID, cartID).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "order_id", "product_id", "name", "description", "photo", "price"}).
							AddRow("3300", orderID, "2020", "iPhone 11", "Apple Smartphone", "", "13000000").
							AddRow("3330", orderID, "2222", "Samsung Galaxy S20 Ultra", "Samsung Smartphone", "", "20000000").
							AddRow("3333", orderID, "2000", "iPhone X", "Apple Smartphone", "", "10000000"),
					)
				mock.ExpectExec("DELETE FROM cart_items WHERE cart_id = $1").
					WithArgs(cartID).
					WillReturnError(sql.ErrTxDone)
			},
			wantOrder: entity.Order{},
			wantErr:   sql.ErrTxDone,
		},
		{
			name:             "error when delete from carts",
			generatedOrderID: orderID,
			cartID:           cartID,
			customerName:     "John Doe",
			customerEmail:    "john.doe@example.com",
			customerAddress:  "1 Hacker Way",
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO orders (id, customer_name, customer_email, customer_address) VALUES ($1, $2, $3, $4) RETURNING id").
					WithArgs(orderID, "John Doe", "john.doe@example.com", "1 Hacker Way").
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).
							AddRow(orderID),
					)
				mock.ExpectQuery(`INSERT INTO order_items (id, order_id, product_id, name, description, photo, price)
				SELECT ci.id, $1, ci.product_id, p.name, p.description, p.photo, p.price
				FROM cart_items ci
				JOIN products p ON p.id = ci.product_id
				WHERE ci.cart_id = $2
				RETURNING id, order_id, product_id, name, description, photo, price`).
					WithArgs(orderID, cartID).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "order_id", "product_id", "name", "description", "photo", "price"}).
							AddRow("3300", orderID, "2020", "iPhone 11", "Apple Smartphone", "", "13000000").
							AddRow("3330", orderID, "2222", "Samsung Galaxy S20 Ultra", "Samsung Smartphone", "", "20000000").
							AddRow("3333", orderID, "2000", "iPhone X", "Apple Smartphone", "", "10000000"),
					)
				mock.ExpectExec("DELETE FROM cart_items WHERE cart_id = $1").
					WithArgs(cartID).
					WillReturnResult(sqlmock.NewResult(0, 3))
				mock.ExpectExec("DELETE FROM carts WHERE cart_id = $1").
					WithArgs(cartID).
					WillReturnError(sql.ErrTxDone)
			},
			wantOrder: entity.Order{},
			wantErr:   sql.ErrTxDone,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := newDBMock()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			if tc.dbMockFn != nil {
				tc.dbMockFn(mock)
			}

			repo := NewOrder(db, uuidGeneratorMock(tc.generatedOrderID))
			gotOrder, gotErr := repo.CheckoutFromCart(tc.cartID, tc.customerName, tc.customerEmail, tc.customerAddress)

			assert.Equal(t, tc.wantErr, gotErr)
			assert.Equal(t, tc.wantOrder, gotOrder)
		})
	}
}
