package cart

import (
	_cart "sepatuku-project/entity/cart"
)

type CartInterfaceRepository interface {
	CreateCart(id int, cart _cart.Cart) (_cart.Cart, error)
	DeleteCart(id int) (_cart.Cart, error)
	GetAllCart() (cart []_cart.Cart, err error)
}
