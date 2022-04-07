package main

import (
	"fmt"
	"log"
	_configs "sepatuku-project/configs"
	_authHandler "sepatuku-project/delivery/handler/auth"
	_cartHadler "sepatuku-project/delivery/handler/cart"
	_orderHandler "sepatuku-project/delivery/handler/order"
	_productHandler "sepatuku-project/delivery/handler/product"
	_userHandler "sepatuku-project/delivery/handler/user"
	_middlewares "sepatuku-project/delivery/middleware"
	_routes "sepatuku-project/delivery/routes"
	_driver "sepatuku-project/driver"
	_authRepository "sepatuku-project/repository/auth"
	_cartRepo "sepatuku-project/repository/cart"
	_orderRepo "sepatuku-project/repository/order"
	_productRepository "sepatuku-project/repository/product"
	_userRepository "sepatuku-project/repository/user"
	_authService "sepatuku-project/service/auth"
	_cartService "sepatuku-project/service/cart"
	_orderService "sepatuku-project/service/order"
	_productService "sepatuku-project/service/product"
	_userService "sepatuku-project/service/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs := _configs.GetConfig()
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

	cartRepo := _cartRepo.NewCartRepository(db)
	cartService := _cartService.NewProductService(cartRepo)
	cartHandler := _cartHadler.NewProductHandler(cartService)

	orderRepo := _orderRepo.NewRepositoryOrder(db)
	orderService := _orderService.NewOrderService(orderRepo)
	orderHandler := _orderHandler.NewOrderHandler(orderService)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomCors())
	e.Use(_middlewares.CustomLogger())

	_routes.UserPath(e, userHandler)
	_routes.ProductPath(e, productHandler)
	_routes.RegisterAuthPath(e, authHandler)
	_routes.CartPath(e, cartHandler)
	_routes.OrderPath(e, orderHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))
}
