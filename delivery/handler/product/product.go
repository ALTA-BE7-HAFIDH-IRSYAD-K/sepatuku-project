package product

import (
	"net/http"
	_middlewares "sepatuku-project/delivery/middleware"
	"sepatuku-project/delivery/response"
	_entities "sepatuku-project/entity/product"
	_poductService "sepatuku-project/service/product"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService _poductService.ProductServiceInterface
}

func NewProductHandler(productService _poductService.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (ph *ProductHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		product, err := ph.productService.GetAllProduct()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed fetch data"))
		}
		if len(product) == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all data", product))
	}
}

func (ph *ProductHandler) GetProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		product, rows, err := ph.productService.GetProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		if product.UserID != uint(idToken) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get data", product))
	}
}

func (ph *ProductHandler) DeleteProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		product, rows, err := ph.productService.GetProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		if idToken != int(product.UserID) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		_, err = ph.productService.DeleteProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed delete data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("Succes delete data"))
	}
}

func (ph *ProductHandler) CreateProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var product _entities.Product
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		c.Bind(&product)
		if idToken != int(product.UserID) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		product, err := ph.productService.CreateProduct(product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Succes create data", product))
	}
}

func (ph *ProductHandler) UpdateProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		product, rows, err := ph.productService.GetProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		if idToken != int(product.UserID) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		c.Bind(&product)
		product, err = ph.productService.UpdateProduct(product, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Success edit data", product))
	}
}
