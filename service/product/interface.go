package product

import _entities "sepatuku-project/entity/product"

type ProductServiceInterface interface {
	GetAllProduct() ([]_entities.Product, error)
	GetProduct(id int) (_entities.Product, int, error)
	DeleteProduct(id int) (_entities.Product, error)
	CreateProduct(product _entities.Product) (_entities.Product, int, error)
	UpdateProduct(product _entities.Product, id int) (_entities.Product, error)
}
