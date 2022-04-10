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
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed get all product"))
		}
		if len(product) == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("product not exist"))
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
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed get product"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("product not exist"))
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
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed get product"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("product not exist"))
		}
		if idToken != int(product.UserID) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("product not exist"))
		}
		_, err = ph.productService.DeleteProduct(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed delete product"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success delete product"))
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
		product, row, err := ph.productService.CreateProduct(product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed create product"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("please complete data filling"))

		}
		if row == 2 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("failed create product"))

		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success create product", product))
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
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed get product"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("product not exist"))
		}
		if idToken != int(product.UserID) {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("product not exist"))
		}
		name_product := product.Name_product
		description := product.Description
		price := product.Price
		image := product.Image
		stock := product.Stock
		c.Bind(&product)
		if product.Name_product == "" {
			product.Name_product = name_product
		}
		if product.Description == "" {
			product.Description = description
		}
		if product.Price == 0 {
			product.Price = price
		}
		if product.Image == "" {
			product.Image = image
		}
		if product.Stock == 0 {
			product.Stock = stock
		}
		product, err = ph.productService.UpdateProduct(product, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed edit product"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success edit product", product))
	}
}
