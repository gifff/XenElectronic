package postgresql

import (
	"github.com/gifff/xenelectronic/contract"
	"github.com/gifff/xenelectronic/entity"
	"github.com/jmoiron/sqlx"
)

func NewCategory(db *sqlx.DB) contract.CategoryRepository {
	return &categoryRepository{db}
}

type categoryRepository struct {
	db *sqlx.DB
}

func (repo *categoryRepository) ListAll() ([]entity.Category, error) {
	rows, err := repo.db.Queryx("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}

	categories := []entity.Category{}
	for rows.Next() {
		category := entity.Category{}
		err = rows.StructScan(&category)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (repo *categoryRepository) ListProductsByCategoryID(categoryID, since int64, limit int32) ([]entity.Product, error) {
	rows, err := repo.db.Queryx(
		"SELECT id, category_id, name, description, photo, price FROM products WHERE category_id = $1 LIMIT $2 OFFSET $3",
		categoryID,
		limit,
		since,
	)
	if err != nil {
		return nil, err
	}

	products := []entity.Product{}
	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.CategoryID, &product.Name, &product.Description, &product.Photo, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
