package order

import (
	"github.com/labstack/echo/v4"
	"net/http"
	_middlewares "sepatuku-project/delivery/middleware"
	"sepatuku-project/delivery/response"
	"sepatuku-project/entity/order"
	_order "sepatuku-project/service/order"
)

type OrderHandler struct {
	serviceOrder _order.ServiceOrderInterface
}

func NewOrderHandler(serviceOrder _order.ServiceOrderInterface) *OrderHandler {
	return &OrderHandler{serviceOrder}
}

func (oh *OrderHandler) GetAllHistoryOrderProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderHistory, err := oh.serviceOrder.GetOrderHistory()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed fetch data"))
		}
		if len(orderHistory) == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all order history", orderHistory))
	}
}

func (oh *OrderHandler) CreateOrderProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newOrder order.Order

		idToken, errToken := _middlewares.ReadTokenId(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}

		if idToken != int(newOrder.UserId) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("user id not found"))
		}

		c.Bind(&newOrder)

		newCart, err := oh.serviceOrder.CreateOrder(idToken, newOrder)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Succes create data", newCart))
	}
}
