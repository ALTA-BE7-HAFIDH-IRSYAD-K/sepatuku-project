package cart

import (
	"sepatuku-project/entity/cart"
	_cart "sepatuku-project/repository/cart"
)

type CartService struct {
	cartService _cart.CartInterfaceRepository
}

func NewProductService(cartRepo _cart.CartInterfaceRepository) CartInterfaceService {
	return &CartService{
		cartService: cartRepo,
	}
}

func (cs *CartService) CreateCart(id int, cart cart.Cart) (cart.Cart, error) {
	//TODO implement me
	cartCreate, err := cs.cartService.CreateCart(id, cart)
	return cartCreate, err

}

func (cs *CartService) DeleteCart(id int) (cart.Cart, error) {
	//TODO implement me
	deleteCart, err := cs.cartService.DeleteCart(id)
	return deleteCart, err
}

func (cs *CartService) GetAllCart() (cart []cart.Cart, err error) {
	//TODO implement me
	getCart, err := cs.cartService.GetAllCart()
	return getCart, err
}
