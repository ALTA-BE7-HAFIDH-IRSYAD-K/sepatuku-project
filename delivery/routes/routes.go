package routes

import (
	_authHandler "sepatuku-project/delivery/handler/auth"
	_cartHandler "sepatuku-project/delivery/handler/cart"
	_orderHandler "sepatuku-project/delivery/handler/order"
	_productHandler "sepatuku-project/delivery/handler/product"
	_userHandler "sepatuku-project/delivery/handler/user"
	_middlewares "sepatuku-project/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/api/v1/auth", ah.LoginHandler())
}

func UserPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/api/v1/users", uh.CreateUserHandler())
	e.GET("/api/v1/users", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/api/v1/users", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/api/v1/users", uh.UpdatedUserHandler(), _middlewares.JWTMiddleware())
}

func CartPath(e *echo.Echo, ch *_cartHandler.CartHandler) {
	e.POST("/api/v1/carts", ch.CreateCartHandler(), _middlewares.JWTMiddleware())
	e.GET("/api/v1/carts", ch.GetAllCartHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/api/v1/carts/:id", ch.DeleteCartHandler(), _middlewares.JWTMiddleware())
	e.GET("/api/v1/carts/:id", ch.GetCartByIdHandler(), _middlewares.JWTMiddleware())
	e.PUT("/api/v1/carts/:id", ch.UpdateQuantityHandler(), _middlewares.JWTMiddleware())
}

func OrderPath(e *echo.Echo, oh *_orderHandler.OrderHandler) {
	e.POST("/api/v1/orders", oh.CreateOrderProduct(), _middlewares.JWTMiddleware())
	e.GET("/api/v1/orders/history", oh.GetAllHistoryOrderProduct(), _middlewares.JWTMiddleware())
	e.PUT("/api/v1/orders/history/status/:id", oh.UpdateOrderStatus(), _middlewares.JWTMiddleware())
}

func ProductPath(e *echo.Echo, ph *_productHandler.ProductHandler) {
	e.POST("/api/v1/products", ph.CreateProductHandler(), _middlewares.JWTMiddleware())
	e.GET("/api/v1/products", ph.GetAllHandler())
	e.GET("/api/v1/products/:id", ph.GetProductHandler())
	e.DELETE("/api/v1/products/:id", ph.DeleteProductHandler(), _middlewares.JWTMiddleware())
	e.PUT("/api/v1/products/:id", ph.UpdateProductHandler(), _middlewares.JWTMiddleware())
}
