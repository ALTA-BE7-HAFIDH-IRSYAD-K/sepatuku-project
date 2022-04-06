package routes

import (
	"fmt"
	_authHandler "sepatuku-project/delivery/handler/auth"
	_productHandler "sepatuku-project/delivery/handler/product"
	_userHandler "sepatuku-project/delivery/handler/user"
	_middlewares "sepatuku-project/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}
func UserPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users", uh.UpdatedUserHandler(), _middlewares.JWTMiddleware())
}

func CartPath() {
	fmt.Println("cart-path")
}

func OrderPath() {
	fmt.Println("order-path")
}

func ProductPath(e *echo.Echo, ph *_productHandler.ProductHandler) {
	e.POST("/products", ph.CreateProductHandler(), _middlewares.JWTMiddleware())
	e.GET("/products", ph.GetAllHandler(), _middlewares.JWTMiddleware())
	e.GET("/products/:id", ph.GetProductHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/products/:id", ph.DeleteProductHandler(), _middlewares.JWTMiddleware())
	e.PUT("/products/:id", ph.UpdateProductHandler(), _middlewares.JWTMiddleware())
}
