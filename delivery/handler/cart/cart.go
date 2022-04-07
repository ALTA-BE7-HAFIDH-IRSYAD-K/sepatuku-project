package cart

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	_middlewares "sepatuku-project/delivery/middleware"
	"sepatuku-project/delivery/response"
	_cart "sepatuku-project/entity/cart"
	"sepatuku-project/service/cart"
	"strconv"
)

type CartHandler struct {
	cartService cart.CartInterfaceService
}

func NewCartHandler(cartService cart.CartInterfaceService) *CartHandler {
	return &CartHandler{cartService}
}

func (ch *CartHandler) GetAllCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		carts, err := ch.cartService.GetAllCart()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed fetch data"))
		}
		if len(carts) == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all data", carts))
	}
}

func (ch *CartHandler) DeleteCartHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var cartDelete _cart.Cart
		idToken, errToken := _middlewares.ReadTokenId(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}

		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)

		if idToken != int(cartDelete.UserId) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
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

		fmt.Println("cartNew", cartNew)

		c.Bind(&cartNew)
		newCart, err := ch.cartService.CreateCart(idToken, cartNew)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Succes create data", newCart))
	}
}
