package main

import (
	"fmt"
	"log"
	"net/http"
	"sepatuku-project/configs"
	_authHandler "sepatuku-project/delivery/handler/auth"
	"sepatuku-project/delivery/handler/cart"
	"sepatuku-project/delivery/handler/order"
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

	cartRepo := _cartRepo.NewCartRepository(db)
	cartService := _cartService.NewCartService(cartRepo)
	cartHandler := cart.NewCartHandler(cartService)

	orderRepo := _orderRepo.NewRepositoryOrder(db)
	orderService := _orderService.NewOrderService(orderRepo)
	orderHandler := order.NewOrderHandler(orderService)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	_routes.UserPath(e, userHandler)
	_routes.ProductPath(e, productHandler)
	_routes.RegisterAuthPath(e, authHandler)
	_routes.CartPath(e, cartHandler)
	_routes.OrderPath(e, orderHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", configs.Port)))
}
