package routes

import (
	"fmt"
	_authHandler "sepatuku-project/delivery/handler/auth"
	_userHandler "sepatuku-project/delivery/handler/user"
	_middlewares "sepatuku-project/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}
func UserPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdatedUserHandler(), _middlewares.JWTMiddleware())
}

func CartPath() {
	fmt.Println("cart-path")
}

func OrderPath() {
	fmt.Println("order-path")
}

func ProductPath() {
	fmt.Println("product-path")
}
