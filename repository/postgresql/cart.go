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
	cartItems := []entity.CartItem{}

	rows, err := repo.db.Queryx(`SELECT ci.id, ci.cart_id, p.id, p.category_id, p.name, p.description, p.photo, p.price FROM products p
	JOIN cart_items ci ON p.id = ci.product_id
	WHERE ci.cart_id = $1`, cartID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cartItem := entity.CartItem{}
		err := rows.Scan(
			&cartItem.ID,
			&cartItem.CartID,
			&cartItem.Product.ID,
			&cartItem.Product.CategoryID,
			&cartItem.Product.Name,
			&cartItem.Product.Description,
			&cartItem.Product.Photo,
			&cartItem.Product.Price,
		)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
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
	var id int64

	row := repo.db.QueryRowx("DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2 RETURNING id", cartID, productID)
	err := row.Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
