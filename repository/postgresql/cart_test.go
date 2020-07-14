package postgresql

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
