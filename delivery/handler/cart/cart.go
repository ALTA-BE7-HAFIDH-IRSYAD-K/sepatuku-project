package cart

import (
	"fmt"
	"net/http"
	_middlewares "sepatuku-project/delivery/middleware"
	"sepatuku-project/delivery/response"
	_cart "sepatuku-project/entity/cart"
	"sepatuku-project/service/cart"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartService cart.CartInterfaceService
}

func NewCartHandler(cartService cart.CartInterfaceService) *CartHandler {
	return &CartHandler{cartService}
}

func (ch *CartHandler) GetAllCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ReadTokenId(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("token is not valid"))
		}

		allCart, _, err := ch.cartService.GetAllCart(idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed fetch data"))
		}

		if len(allCart) == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all data allCart", allCart))
	}
}

func (ch *CartHandler) DeleteCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//var cartDelete _cart.Cart
		_, errToken := _middlewares.ReadTokenId(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("token is not valid"))
		}

		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)

		//if idToken != int(cartDelete.UserId) {
		//	return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		//}

		_, err := ch.cartService.DeleteCart(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed delete data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("Success delete data"))
	}
}

func (ch *CartHandler) CreateCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var cartNew _cart.Cart

		idToken, errToken := _middlewares.ReadTokenId(c)
		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}

		cartNew.UserId = uint(idToken)

		fmt.Println("cartNew", cartNew)

		c.Bind(&cartNew)

		newCart, err := ch.cartService.CreateCart(cartNew)

		//for _, value := range product {
		//	value.ID = newCart.Product.ID
		//	value.Name_product = newCart.Product.Name_product
		//	value.Description = newCart.Product.Description
		//	value.Price = newCart.Product.Price
		//	value.Image = newCart.Product.Image
		//	value.Stock = newCart.Product.Stock
		//}

		fmt.Println("newCart", newCart)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success create data cart", newCart))
	}
}
