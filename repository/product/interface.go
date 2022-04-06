package product

import _entities "sepatuku-project/entity/product"

type ProductRepositoryInterface interface {
	GetAllProduct() ([]_entities.Product, error)
	GetProduct(id int) (_entities.Product, int, error)
	DeleteProduct(id int) (_entities.Product, error)
	CreateProduct(idToken int, product _entities.Product) (_entities.Product, error)
	UpdateProduct(product _entities.Product) (_entities.Product, error)
}
