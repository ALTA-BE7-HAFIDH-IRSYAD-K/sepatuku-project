package cart

import (
	"fmt"
	_entities "sepatuku-project/entity/cart"
	"sepatuku-project/entity/product"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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
	var cartObject []_entities.CartResponse
	var productId []product.Product

	tx := cr.database.Joins("left join products on carts.product_id = products.id").Joins("left join users on carts.user_id = users.id").Where("carts.user_id", id).Find(&cart)

	copier.Copy(&cartObject, &cart)

	fmt.Println(cartObject, "cartObject")

	if tx.Error != nil {
		return cart, productId, tx.Error
	}

	return cart, productId, nil
}

func (cr *CartRepository) CreateCart(cart _entities.Cart) (_entities.CartResponseCreate, error) {
	var cartObject _entities.CartResponseCreate

	fmt.Println(cart, "Add-Cart")

	err := cr.database.Where("id", cart.ProductId).Find(&cart)

	tx := cr.database.Save(&cart)

	copier.Copy(&cartObject, &cart)

	if tx.Error != nil {
		return cartObject, err.Error
	}

	if tx.Error != nil {
		return cartObject, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cartObject, tx.Error
	}

	fmt.Println("cartObject", cartObject)

	return cartObject, nil
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

func (cr *CartRepository) GetCartById(id int) (_entities.Cart, int, error) {
	var cart _entities.Cart
	tx := cr.database.Find(&cart, id)

	if tx.Error != nil {
		return cart, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return cart, 0, tx.Error
	}
	return cart, int(tx.RowsAffected), nil
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
