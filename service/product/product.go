package product

import (
	_entities "sepatuku-project/entity/product"
	_productRepository "sepatuku-project/repository/product"
)

type ProductService struct {
	productRepository _productRepository.ProductRepositoryInterface
}

func NewProductService(productRepo _productRepository.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{
		productRepository: productRepo,
	}
}

func (ps *ProductService) GetAllProduct() ([]_entities.Product, error) {
	product, err := ps.productRepository.GetAllProduct()
	return product, err
}

func (ps *ProductService) GetProduct(id int) (_entities.Product, int, error) {
	product, rows, err := ps.productRepository.GetProduct(id)
	return product, rows, err
}

func (ps *ProductService) DeleteProduct(id int) (_entities.Product, error) {
	product, err := ps.productRepository.DeleteProduct(id)
	return product, err
}
func (ps *ProductService) CreateProduct(product _entities.Product) (_entities.Product, error) {
	product, err := ps.productRepository.CreateProduct(product)
	return product, err
}
func (ps *ProductService) UpdateProduct(product _entities.Product, id int) (_entities.Product, error) {
	product.ID = uint(id)
	product, err := ps.productRepository.UpdateProduct(product)
	return product, err
}
