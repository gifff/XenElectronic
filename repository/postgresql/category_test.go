package postgresql

import (
	"database/sql"
	"testing"

	"github.com/gifff/xenelectronic/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func newDBMock() (*sqlx.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, nil, err
	}

	dbx := sqlx.NewDb(db, "sqlmock")
	return dbx, mock, nil
}

func TestListAll(t *testing.T) {
	db, mock, err := newDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, name FROM categories").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1000, "Home Appliances").
			AddRow(1001, "Handphones").
			AddRow(1010, "Smartphones"),
	)

	repo := NewCategory(db)
	gotCategories, gotErr := repo.ListAll()

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

func TestListAll_Error(t *testing.T) {
	db, mock, err := newDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, name FROM categories").WillReturnError(sql.ErrConnDone)

	repo := NewCategory(db)
	gotCategories, gotErr := repo.ListAll()

	assert.Equal(t, sql.ErrConnDone, gotErr)
	assert.Nil(t, gotCategories)
}

func TestListProductsByCategoryID(t *testing.T) {
	db, mock, err := newDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, category_id, name, description, photo, price FROM products WHERE category_id = $1 LIMIT $2 OFFSET $3").
		WithArgs(int64(1010), int32(5), int64(10)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "category_id", "name", "description", "photo", "price"}).
				AddRow(2000, 1010, "iPhone X", "Apple Smartphone", "", 10000000).
				AddRow(2001, 1010, "iPhone XR", "Apple Smartphone", "", 11000000).
				AddRow(2010, 1010, "iPhone XS", "Apple Smartphone", "", 15000000).
				AddRow(2020, 1010, "iPhone 11", "Apple Smartphone", "", 13000000).
				AddRow(2222, 1010, "Samsung Galaxy S20 Ultra", "Samsung Smartphone", "", 20000000),
		)

	repo := NewCategory(db)
	gotProducts, gotErr := repo.ListProductsByCategoryID(1010, 10, 5)

	wantProducts := []entity.Product{
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
	assert.Equal(t, wantProducts, gotProducts)
}

func TestListProductsByCategoryID_Error(t *testing.T) {
	db, mock, err := newDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, category_id, name, description, photo, price FROM products WHERE category_id = $1 LIMIT $2 OFFSET $3").
		WithArgs(int64(1010), int32(5), int64(10)).
		WillReturnError(sql.ErrConnDone)

	repo := NewCategory(db)
	gotProducts, gotErr := repo.ListProductsByCategoryID(1010, 10, 5)

	assert.Equal(t, sql.ErrConnDone, gotErr)
	assert.Nil(t, gotProducts)
}
