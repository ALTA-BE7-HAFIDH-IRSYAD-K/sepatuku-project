package cart

import (
	"sepatuku-project/entity/cart"
	"sepatuku-project/entity/product"
	_cart "sepatuku-project/repository/cart"
)

type CartService struct {
	cartService _cart.CartInterfaceRepository
}

func NewCartService(cartRepo _cart.CartInterfaceRepository) CartInterfaceService {
	return &CartService{
		cartService: cartRepo,
	}
}

func (cs *CartService) CreateCart(cart cart.Cart) (cart.Cart, error) {
	//TODO implement me
	cartCreate, err := cs.cartService.CreateCart(cart)
	return cartCreate, err
}

func (cs *CartService) DeleteCart(id int) (cart.Cart, error) {
	//TODO implement me
	deleteCart, err := cs.cartService.DeleteCart(id)
	return deleteCart, err
}

func (cs *CartService) GetAllCart(id int) ([]cart.Cart, []product.Product, error) {
	//TODO implement me
	cart, product, err := cs.cartService.GetAllCart(id)
	return cart, product, err
}
