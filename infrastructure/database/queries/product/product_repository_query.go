package storage

import (
	"database/sql"
	entity "github.com/StuardAP/clean_rest_golang/internal/domain/entities"
)

type ProductRepositoryMySQL struct {
	db *sql.DB
}

func (r *ProductRepositoryMySQL) CreateProduct(product *entity.Product) error {
	_, err := r.db.Exec("INSERT INTO products (id, name, price, created_at) VALUES (?, ?, ?, ?)",
		product.ID, product.Name, product.Price, product.CreatedAt)
	return err
}

func (r *ProductRepositoryMySQL) GetProductByID(id string) (*entity.Product, error) {
	panic("unimplemented")
}

func NewProductRepositoryMySQL(db *sql.DB) *ProductRepositoryMySQL {
	return &ProductRepositoryMySQL{db: db}
}

func (r *ProductRepositoryMySQL) GetAllProducts() ([]entity.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price, created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
