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
	cartItem := entity.CartItem{
		CartID: cartID,
	}

	row := repo.db.QueryRowx("INSERT INTO cart_items(cart_id, product_id) VALUES($1, $2) RETURNING id", cartID, productID)
	err := row.Scan(&cartItem.ID)
	if err != nil {
		return entity.CartItem{}, err
	}

	row = repo.db.QueryRowx("SELECT id, category_id, name, description, photo, price FROM products WHERE id = $1", productID)
	err = row.Scan(
		&cartItem.Product.ID,
		&cartItem.Product.CategoryID,
		&cartItem.Product.Name,
		&cartItem.Product.Description,
		&cartItem.Product.Photo,
		&cartItem.Product.Price,
	)
	if err != nil {
		return entity.CartItem{}, err
	}

	return cartItem, nil
}

func (repo *cartRepository) RemoveProductFromCart(cartID string, productID int64) error {
	return nil
}
