package order

import (
	"fmt"
	"net/http"
	_middlewares "sepatuku-project/delivery/middleware"
	"sepatuku-project/delivery/response"
	"sepatuku-project/entity/order"
	_order "sepatuku-project/service/order"
	"strconv"

	"github.com/labstack/echo/v4"
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
		//var product _product.Product

		idToken, errToken := _middlewares.ReadTokenId(c)

		newOrder.UserId = uint(idToken)
		//newOrder.ProductId = product.ID

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}

		if idToken != int(newOrder.UserId) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("user id not found"))
		}

		//if newOrder.ProductId != product.ID {
		//	return c.JSON(http.StatusBadRequest, response.ResponseFailed("product id not found"))
		//}
		fmt.Println("newOrder", newOrder)

		c.Bind(&newOrder)

		fmt.Println("newOrder", newOrder)

		newCart, err := oh.serviceOrder.CreateOrder(newOrder)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Succes create data", newCart))
	}
}

func (oh *OrderHandler) UpdateOrderStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var orderUpdate order.Order

		idToken, tokenErr := _middlewares.ReadTokenId(c)
		orderUpdate.UserId = uint(idToken)

		if tokenErr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)

		c.Bind(&orderUpdate)

		orderId, err := oh.serviceOrder.UpdatedHistoryOrder(orderUpdate, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Success edit status", orderId))
	}
}

func (oh *OrderHandler) GetOrderByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenErr := _middlewares.ReadTokenId(c)
		if tokenErr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		orderById, rows, err := oh.serviceOrder.GetOrderHistoryById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		if orderById.UserId != uint(idToken) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get data", orderById))
	}
}
