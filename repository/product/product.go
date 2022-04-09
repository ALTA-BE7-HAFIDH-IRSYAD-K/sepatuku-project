package product

import (
	"fmt"
	_entities "sepatuku-project/entity/product"

	"gorm.io/gorm"
)

type ProductRepository struct {
	database *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		database: db,
	}
}

func (pr *ProductRepository) GetAllProduct() ([]_entities.Product, error) {
	var product []_entities.Product
	tx := pr.database.Find(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product, nil
}

func (pr *ProductRepository) GetProduct(id int) (_entities.Product, int, error) {
	var product _entities.Product
	tx := pr.database.Find(&product, id)
	if tx.Error != nil {
		return product, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return product, 0, tx.Error
	}
	return product, int(tx.RowsAffected), nil
}

func (pr *ProductRepository) CreateProduct(product _entities.Product) (_entities.Product, int, error) {
	if product.Name_product == "" {
		return product, 0, fmt.Errorf("please complete data filling")
	}
	if product.Description == "" {
		return product, 0, fmt.Errorf("please complete data filling")
	}
	if product.Price == 0 {
		return product, 0, fmt.Errorf("please complete data filling")
	}
	if product.Image == "" {
		return product, 0, fmt.Errorf("please complete data filling")
	}
	if product.Stock == 0 {
		return product, 0, fmt.Errorf("please complete data filling")
	}
	tx := pr.database.Save(&product)
	if tx.Error != nil {
		return product, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return product, 2, tx.Error
	}
	return product, int(tx.RowsAffected), nil
}

func (pr *ProductRepository) UpdateProduct(product _entities.Product) (_entities.Product, error) {
	tx := pr.database.Save(&product)
	if tx.Error != nil {
		return product, tx.Error
	}
	if tx.RowsAffected == 0 {
		return product, tx.Error
	}
	return product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) (_entities.Product, error) {
	var product _entities.Product
	tx := pr.database.Delete(&product, id)
	if tx.Error != nil {
		return product, tx.Error
	}
	if tx.RowsAffected == 0 {
		return product, tx.Error

	}
	return product, nil
}
