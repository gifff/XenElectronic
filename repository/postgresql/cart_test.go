package postgresql

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gifff/xenelectronic/entity"
	"github.com/stretchr/testify/assert"
)

type uuidGeneratorMock string

func (u uuidGeneratorMock) GenerateV4() string {
	return string(u)
}

func TestCreateCart(t *testing.T) {
	db, mock, err := newDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	u := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	mock.ExpectQuery("INSERT INTO carts(cart_id) VALUES($1) RETURNING cart_id").
		WithArgs(u).
		WillReturnRows(
			sqlmock.NewRows([]string{"cart_id"}).
				AddRow(u),
		)
	repo := NewCart(db, uuidGeneratorMock(u))
	gotCartID, gotErr := repo.CreateCart()

	wantCartID := u

	assert.Nil(t, gotErr)
	assert.Equal(t, wantCartID, gotCartID)
}

func TestCreateCart_Error(t *testing.T) {
	db, mock, err := newDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	u := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	mock.ExpectQuery("INSERT INTO carts(cart_id) VALUES($1) RETURNING cart_id").
		WithArgs(u).
		WillReturnError(sql.ErrConnDone)
	repo := NewCart(db, uuidGeneratorMock(u))
	gotCartID, gotErr := repo.CreateCart()

	wantCartID := ""
	wantErr := sql.ErrConnDone

	assert.Equal(t, wantErr, gotErr)
	assert.Equal(t, wantCartID, gotCartID)
}

func TestAddProductIntoCart(t *testing.T) {
	u := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	testCases := []struct {
		name         string
		cartID       string
		productID    int64
		dbMockFn     func(mock sqlmock.Sqlmock)
		wantCartItem entity.CartItem
		wantErr      error
	}{
		{
			name:      "success",
			cartID:    u,
			productID: 2001,
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO cart_items(cart_id, product_id) VALUES($1, $2) RETURNING id").
					WithArgs(u, int64(2001)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).
							AddRow(3456),
					)
				mock.ExpectQuery("SELECT id, category_id, name, description, photo, price FROM products WHERE id = $1").
					WithArgs(int64(2001)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "category_id", "name", "description", "photo", "price"}).
							AddRow(2001, 1010, "iPhone XR", "Apple Smartphone", "", 11000000),
					)
			},
			wantCartItem: entity.CartItem{
				ID:     3456,
				CartID: u,
				Product: entity.Product{
					ID:          2001,
					CategoryID:  1010,
					Name:        "iPhone XR",
					Description: "Apple Smartphone",
					Price:       11000000,
				},
			},
		},
		{
			name:      "error when insert into cart items",
			cartID:    u,
			productID: 2001,
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO cart_items(cart_id, product_id) VALUES($1, $2) RETURNING id").
					WithArgs(u, int64(2001)).
					WillReturnError(sql.ErrConnDone)
			},
			wantCartItem: entity.CartItem{},
			wantErr:      sql.ErrConnDone,
		},
		{
			name:      "error when select from products",
			cartID:    u,
			productID: 2001,
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO cart_items(cart_id, product_id) VALUES($1, $2) RETURNING id").
					WithArgs(u, int64(2001)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).
							AddRow(3456),
					)
				mock.ExpectQuery("SELECT id, category_id, name, description, photo, price FROM products WHERE id = $1").
					WithArgs(int64(2001)).
					WillReturnError(sql.ErrConnDone)
			},
			wantCartItem: entity.CartItem{},
			wantErr:      sql.ErrConnDone,
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

			repo := NewCart(db, nil)
			gotCartItem, gotErr := repo.AddProductIntoCart(tc.cartID, tc.productID)

			assert.Equal(t, tc.wantErr, gotErr)
			assert.Equal(t, tc.wantCartItem, gotCartItem)
		})
	}
}

func TestListProductsByCartID(t *testing.T) {
	db, mock, err := newDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	u := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	mock.ExpectQuery(`SELECT ci.id, ci.cart_id, p.id, p.category_id, p.name, p.description, p.photo, p.price FROM products p
	JOIN cart_items ci ON p.id = ci.product_id
	WHERE ci.cart_id = $1`).
		WithArgs(u).
		WillReturnRows(
			sqlmock.NewRows([]string{"ci.id", "ci.cart_id", "p.id", "p.category_id", "p.name", "p.description", "p.photo", "p.price"}).
				AddRow(3456, u, 2001, 1010, "iPhone XR", "Apple Smartphone", "", 11000000).
				AddRow(3457, u, 2010, 1010, "iPhone XS", "Apple Smartphone", "", 15000000).
				AddRow(3458, u, 2020, 1010, "iPhone 11", "Apple Smartphone", "", 13000000),
		)
	repo := NewCart(db, nil)
	gotCartItems, gotErr := repo.ListProductsByCartID(u)

	wantCartItems := []entity.CartItem{
		{
			ID:     3456,
			CartID: u,
			Product: entity.Product{
				ID:          2001,
				CategoryID:  1010,
				Name:        "iPhone XR",
				Description: "Apple Smartphone",
				Price:       11000000,
			},
		},
		{
			ID:     3457,
			CartID: u,
			Product: entity.Product{
				ID:          2010,
				CategoryID:  1010,
				Name:        "iPhone XS",
				Description: "Apple Smartphone",
				Price:       15000000,
			},
		},
		{
			ID:     3458,
			CartID: u,
			Product: entity.Product{
				ID:          2020,
				CategoryID:  1010,
				Name:        "iPhone 11",
				Description: "Apple Smartphone",
				Price:       13000000,
			},
		},
	}

	assert.Nil(t, gotErr)
	assert.Equal(t, wantCartItems, gotCartItems)
}

func TestListProductsByCartID_Error(t *testing.T) {
	db, mock, err := newDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	u := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	mock.ExpectQuery(`SELECT ci.id, ci.cart_id, p.id, p.category_id, p.name, p.description, p.photo, p.price FROM products p
	JOIN cart_items ci ON p.id = ci.product_id
	WHERE ci.cart_id = $1`).
		WithArgs(u).
		WillReturnError(sql.ErrConnDone)
	repo := NewCart(db, nil)
	gotCartItems, gotErr := repo.ListProductsByCartID(u)

	wantCartItems := ([]entity.CartItem)(nil)

	assert.Equal(t, sql.ErrConnDone, gotErr)
	assert.Equal(t, wantCartItems, gotCartItems)
}

func TestRemoveProductFromCart(t *testing.T) {
	u := "aaaaaaaa-bbbb-cccc-dddd-eeeedeadbeef"

	testCases := []struct {
		name      string
		cartID    string
		productID int64
		dbMockFn  func(mock sqlmock.Sqlmock)
		wantErr   error
	}{
		{
			name:      "success",
			cartID:    u,
			productID: 2001,
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2 RETURNING id").
					WithArgs(u, int64(2001)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).
							AddRow(3456),
					)
			},
		},
		{
			name:      "error when delete from cart items",
			cartID:    u,
			productID: 2001,
			dbMockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2 RETURNING id").
					WithArgs(u, int64(2001)).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: sql.ErrNoRows,
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

			repo := NewCart(db, nil)
			gotErr := repo.RemoveProductFromCart(tc.cartID, tc.productID)

			assert.Equal(t, tc.wantErr, gotErr)
		})
	}
}
