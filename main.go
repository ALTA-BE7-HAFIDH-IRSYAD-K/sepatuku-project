package main

import (
	"fmt"
	"log"
	"sepatuku-project/configs"
	_authHandler "sepatuku-project/delivery/handler/auth"
	_productHandler "sepatuku-project/delivery/handler/product"
	_userHandler "sepatuku-project/delivery/handler/user"
	_middlewares "sepatuku-project/delivery/middleware"
	_routes "sepatuku-project/delivery/routes"
	_driver "sepatuku-project/driver"
	_authRepository "sepatuku-project/repository/auth"
	_productRepository "sepatuku-project/repository/product"
	_userRepository "sepatuku-project/repository/user"
	_authService "sepatuku-project/service/auth"
	_productService "sepatuku-project/service/product"
	_userService "sepatuku-project/service/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs := configs.GetConfig()
	db := _driver.InitDB(configs)

	userRepo := _userRepository.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo)
	userHandler := _userHandler.NewUserHandler(userService)
	productRepo := _productRepository.NewProductRepository(db)
	productService := _productService.NewProductService(productRepo)
	productHandler := _productHandler.NewProductHandler(productService)
	authRepo := _authRepository.NewAuthRepository(db)
	authService := _authService.NewAuthService(authRepo)
	authHandler := _authHandler.NewAuthHandler(authService)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())
	// e.Use(_middlewares.CustomCors())
	_routes.UserPath(e, userHandler)
	_routes.ProductPath(e, productHandler)
	_routes.RegisterAuthPath(e, authHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))
}
