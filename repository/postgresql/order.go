package postgresql

import (
	"github.com/gifff/xenelectronic/contract"
	"github.com/gifff/xenelectronic/entity"
	"github.com/jmoiron/sqlx"
)

func NewOrder(db *sqlx.DB, uuidGenerator contract.UUIDGenerator) contract.OrderRepository {
	return &orderRepository{db, uuidGenerator}
}

type orderRepository struct {
	db            *sqlx.DB
	uuidGenerator contract.UUIDGenerator
}

func (repo *orderRepository) CheckoutFromCart(cartID, customerName, customerEmail, customerAddress string) (entity.Order, error) {
	order := entity.Order{
		CustomerName:    customerName,
		CustomerEmail:   customerEmail,
		CustomerAddress: customerAddress,
	}
	orderID := repo.uuidGenerator.GenerateV4()

	tx, err := repo.db.Beginx()
	if err != nil {
		return entity.Order{}, err
	}

	row := tx.QueryRowx("INSERT INTO orders (id, customer_name, customer_email, customer_address) VALUES ($1, $2, $3, $4) RETURNING id",
		orderID,
		customerName,
		customerEmail,
		customerAddress,
	)
	err = row.Scan(&order.ID)
	if err != nil {
		return entity.Order{}, err
	}

	rows, err := tx.Queryx(`INSERT INTO order_items (id, order_id, product_id, name, description, photo, price)
	SELECT ci.id, $1, ci.product_id, p.name, p.description, p.photo, p.price
	FROM cart_items ci
	JOIN products p ON p.id = ci.product_id
	WHERE ci.cart_id = $2
	RETURNING id, order_id, product_id, name, description, photo, price`, orderID, cartID)
	if err != nil {
		return entity.Order{}, err
	}

	for rows.Next() {
		cartItem := entity.CartItem{}
		err := rows.Scan(
			&cartItem.ID,
			&cartItem.CartID,
			&cartItem.Product.ID,
			&cartItem.Product.Name,
			&cartItem.Product.Description,
			&cartItem.Product.Photo,
			&cartItem.Product.Price,
		)
		if err != nil {
			return entity.Order{}, err
		}

		order.CartItems = append(order.CartItems, cartItem)
	}

	_, err = tx.Exec("DELETE FROM cart_items WHERE cart_id = $1", cartID)
	if err != nil {
		return entity.Order{}, err
	}

	_, err = tx.Exec("DELETE FROM carts WHERE cart_id = $1", cartID)
	if err != nil {
		return entity.Order{}, err
	}

	err = tx.Commit()
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (repo *orderRepository) FetchOne(orderID string) (entity.Order, error) {
	return entity.Order{}, nil
}
