package cart

import (
	"fmt"
	"gorm.io/gorm"
	_entities "sepatuku-project/entity/cart"
	"sepatuku-project/entity/product"
)

type CartRepository struct {
	database *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		database: db,
	}
}

func (cr *CartRepository) GetAllCart(id int) ([]_entities.Cart, []product.Product, error) {
	var cart []_entities.Cart
	var productId []product.Product

	idCart := make([]int, 0)

	tx := cr.database.Where("user_id", id).Find(&cart)

	for i := 0; i < len(cart); i++ {
		idCart = append(idCart, int(cart[i].ProductId))
	}

	err := cr.database.Where("id", idCart).Find(&productId)

	fmt.Println("productId", productId)
	fmt.Println("cart", cart)

	if tx.Error != nil {
		return cart, productId, tx.Error
	}

	if tx.Error != nil {
		return cart, productId, err.Error
	}

	return cart, productId, nil
}

func (cr *CartRepository) CreateCart(cart _entities.Cart) (_entities.Cart, error) {
	//var productId product.Product

	fmt.Println(cart, "Add-Cart")

	err := cr.database.Where("id", cart.ProductId).Find(&cart)

	tx := cr.database.Save(&cart)

	if tx.Error != nil {
		return cart, err.Error
	}

	if tx.Error != nil {
		return cart, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cart, tx.Error
	}

	return cart, nil
}

func (cr *CartRepository) DeleteCart(id int) (_entities.Cart, error) {
	var cart _entities.Cart
	tx := cr.database.Delete(id)

	if tx.Error != nil {
		return cart, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cart, tx.Error

	}
	return cart, nil
}

func (cr *CartRepository) UpdateQuantity(cart _entities.Cart) (_entities.Cart, error) {
	tx := cr.database.Save(&cart)
	if tx.Error != nil {
		return cart, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cart, tx.Error
	}
	return cart, nil
}
