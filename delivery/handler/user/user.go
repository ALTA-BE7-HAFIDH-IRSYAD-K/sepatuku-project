package user

import (
	"net/http"
	_middlewares "sepatuku-project/delivery/middleware"
	"sepatuku-project/delivery/response"
	_entities "sepatuku-project/entity/user"
	_userService "sepatuku-project/service/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService _userService.UserServiceInterface
}

func NewUserHandler(userService _userService.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("bad request"))
		}
		id := idToken
		user, product, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to get user profile"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseUser("success get user profile", user, product))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("bad request"))
		}
		id := idToken
		_, _, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		_, err = uh.userService.DeleteUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed delete user"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success delete user"))
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.User
		c.Bind(&user)
		userRes, err := uh.userService.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("register failed"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("register successfully", userRes))
	}
}

func (uh *UserHandler) UpdatedUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("bad request"))
		}
		id := idToken
		user, _, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		c.Bind(&user)
		userRes, err := uh.userService.UpdatedUser(user, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed edit user profile"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success edit user profile", userRes))
	}
}
