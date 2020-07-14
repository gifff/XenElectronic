package postgresql

import (
	"github.com/gifff/xenelectronic/contract"
	"github.com/gifff/xenelectronic/entity"
	"github.com/jmoiron/sqlx"
)

func NewCart(db *sqlx.DB, uuidGenerator contract.UUIDGenerator) contract.CartRepository {
	return &cartRepository{db, uuidGenerator}
}

type cartRepository struct {
	db            *sqlx.DB
	uuidGenerator contract.UUIDGenerator
}

func (repo *cartRepository) CreateCart() (string, error) {
	u := repo.uuidGenerator.GenerateV4()
	row := repo.db.QueryRowx("INSERT INTO carts(cart_id) VALUES($1) RETURNING cart_id", u)
	err := row.Scan(&u)
	if err != nil {
		return "", err
	}

	return u, nil
}

func (repo *cartRepository) ListProductsByCartID(cartID string) ([]entity.CartItem, error) {
	return nil, nil
}

func (repo *cartRepository) AddProductIntoCart(cartID string, productID int64) (entity.CartItem, error) {
	return entity.CartItem{}, nil
}

func (repo *cartRepository) RemoveProductFromCart(cartID string, productID int64) error {
	return nil
}
