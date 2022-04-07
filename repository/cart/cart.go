package cart

import (
	"fmt"
	"gorm.io/gorm"
	_entities "sepatuku-project/entity/cart"
)

type CartRepository struct {
	database *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		database: db,
	}
}

func (cr *CartRepository) GetAllCart() ([]_entities.Cart, error) {
	var cart []_entities.Cart
	tx := cr.database.Find(&cart)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return cart, nil
}

func (cr *CartRepository) CreateCart(id int, cart _entities.Cart) (_entities.Cart, error) {
	//var productId product.Product

	cart.UserId = uint(id)
	//cart.ProductId = productId.ID

	fmt.Println(cart.Product, "cart")

	tx := cr.database.Save(&cart)

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
	tx := cr.database.Delete(&cart, id)

	if tx.Error != nil {
		return cart, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cart, tx.Error

	}
	return cart, nil
}
