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
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all product", product))
	}
}

func (ph *ProductHandler) GetProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		product, rows, err := ph.productService.GetProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get product", product))
	}
}

func (ph *ProductHandler) DeleteProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("bad request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		product, rows, err := ph.productService.GetProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		if idToken != int(product.UserID) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		_, err = ph.productService.DeleteProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed delete data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success delete data"))
	}
}

func (ph *ProductHandler) CreateProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var product _entities.Product
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("bad request"))
		}
		c.Bind(&product)
		product.UserID = uint(idToken)
		product, err := ph.productService.CreateProduct(product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed create data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success create data", product))
	}
}

func (ph *ProductHandler) UpdateProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("bad request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		product, rows, err := ph.productService.GetProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		if idToken != int(product.UserID) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		c.Bind(&product)
		product, err = ph.productService.UpdateProduct(product, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed edit product"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success edit product", product))
	}
}
