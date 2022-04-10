package cart

import (
	_cart "sepatuku-project/entity/cart"
	"sepatuku-project/entity/product"
)

type CartInterfaceRepository interface {
	CreateCart(cart _cart.Cart) (_cart.Cart, error)
	DeleteCart(id int) (_cart.Cart, error)
	//GetAllCart() ([]product.Product, error)
	GetAllCart(id int) ([]_cart.Cart, []product.Product, error)
}
