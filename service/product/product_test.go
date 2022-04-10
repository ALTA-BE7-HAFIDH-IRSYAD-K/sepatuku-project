package product

import (
	"fmt"
	"sepatuku-project/entity/product"
	_entities "sepatuku-project/entity/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProduct(t *testing.T) {
	t.Run("TestGetAllProductSuccess", func(t *testing.T) {
		productService := NewProductService(mockProductRepository{})
		products, err := productService.GetAllProduct()
		assert.Nil(t, err)
		assert.Equal(t, "jordan", products[0].Name_product)
	})

	t.Run("TestGetAllProductError", func(t *testing.T) {
		productService := NewProductService(mockProductRepositoryError{})
		products, err := productService.GetAllProduct()
		assert.NotNil(t, err)
		assert.Equal(t, []product.Product{}, products)
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("TestGetProductSuccess", func(t *testing.T) {
		productService := NewProductService(mockProductRepository{})
		product, row, err := productService.GetProduct(1)
		assert.Nil(t, err)
		assert.Equal(t, "jordan", product.Name_product)
		assert.Equal(t, 1, row)
	})

	t.Run("TestGetProductError", func(t *testing.T) {
		productService := NewProductService(mockProductRepositoryError{})
		products, row, err := productService.GetProduct(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
		assert.Equal(t, product.Product{}, products)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("TestDeleteProductSuccess", func(t *testing.T) {
		productService := NewProductService(mockProductRepository{})
		product, err := productService.DeleteProduct(1)
		assert.Nil(t, err)
		assert.Equal(t, "deleted", product.Name_product)
	})

	t.Run("TestDeleteProductError", func(t *testing.T) {
		productService := NewProductService(mockProductRepositoryError{})
		product, err := productService.DeleteProduct(1)
		assert.NotNil(t, err)
		assert.Equal(t, "jordan", product.Name_product)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("TestCreateProductSuccess", func(t *testing.T) {
		productService := NewProductService(mockProductRepository{})
		products, row, err := productService.CreateProduct(product.Product{
			Name_product: "jordan"})
		assert.Nil(t, err)
		assert.Equal(t, "jordan", products.Name_product)
		assert.Equal(t, 1, row)
	})

	t.Run("TestCreateProductError", func(t *testing.T) {
		productService := NewProductService(mockProductRepositoryError{})
		products, row, err := productService.CreateProduct(product.Product{
			Name_product: "jordan"})
		assert.NotNil(t, err)
		assert.Equal(t, product.Product{}, products)
		assert.Equal(t, 0, row)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("TestUpdateProductSuccess", func(t *testing.T) {
		productService := NewProductService(mockProductRepository{})
		product, err := productService.UpdateProduct(product.Product{
			Name_product: "air"}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "air", product.Name_product)
	})

	t.Run("TestUpdateProductError", func(t *testing.T) {
		productService := NewProductService(mockProductRepositoryError{})
		product, err := productService.UpdateProduct(product.Product{
			Name_product: "air"}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "jordan", product.Name_product)
	})
}

// mock succes
type mockProductRepository struct{}

func (m mockProductRepository) GetAllProduct() ([]_entities.Product, error) {
	return []_entities.Product{{Name_product: "jordan"}}, nil
}

func (m mockProductRepository) GetProduct(id int) (_entities.Product, int, error) {
	return _entities.Product{Name_product: "jordan"}, 1, nil
}

func (m mockProductRepository) DeleteProduct(id int) (_entities.Product, error) {
	return _entities.Product{Name_product: "deleted"}, nil
}

func (m mockProductRepository) UpdateProduct(Product _entities.Product) (_entities.Product, error) {
	return _entities.Product{Name_product: "air"}, nil
}
func (m mockProductRepository) CreateProduct(Product _entities.Product) (_entities.Product, int, error) {
	return _entities.Product{Name_product: "jordan"}, 1, nil
}

//  mock error

type mockProductRepositoryError struct{}

func (m mockProductRepositoryError) GetAllProduct() ([]_entities.Product, error) {
	return []_entities.Product{}, fmt.Errorf("error get Product")
}
func (m mockProductRepositoryError) GetProduct(id int) (_entities.Product, int, error) {
	return _entities.Product{}, 0, fmt.Errorf("error get Product")
}
func (m mockProductRepositoryError) DeleteProduct(id int) (_entities.Product, error) {
	return _entities.Product{Name_product: "jordan"}, fmt.Errorf("error delete Product")
}
func (m mockProductRepositoryError) CreateProduct(_entities.Product) (_entities.Product, int, error) {
	return _entities.Product{}, 0, fmt.Errorf("error create Product")
}
func (m mockProductRepositoryError) UpdateProduct(_entities.Product) (_entities.Product, error) {
	return _entities.Product{Name_product: "jordan"}, fmt.Errorf("error updated Product")
}
