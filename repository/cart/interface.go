package cart

import (
	_cart "sepatuku-project/entity/cart"
	"sepatuku-project/entity/product"
)

type CartInterfaceRepository interface {
	CreateCart(cart _cart.Cart) (_cart.Cart, error)
	DeleteCart(id int) (_cart.Cart, error)
	GetCartById(id int) (_cart.Cart, int, error)
	GetAllCart(id int) ([]_cart.Cart, []product.Product, error)
	UpdateQuantity(cart _cart.Cart) (_cart.Cart, error)
}
