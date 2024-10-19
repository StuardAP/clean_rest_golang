package repositories

import (
    entity "github.com/StuardAP/clean_rest_golang/internal/domain/entities"
)

type ProductRepository interface {
		GetAllProducts() ([]entity.Product, error)
		GetProductByID(id string) (*entity.Product, error)
		CreateProduct(product *entity.Product) error
}

