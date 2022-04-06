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
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		id := idToken
		user, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get data", user))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		id := idToken
		_, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		_, err = uh.userService.DeleteUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed delete data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("Succes delete data"))
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.User
		c.Bind(&user)
		user, err := uh.userService.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Succes create data", user))
	}
}

func (uh *UserHandler) UpdatedUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		id := idToken
		user, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		if idToken != id {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		c.Bind(&user)
		user, err = uh.userService.UpdatedUser(user, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Success edit data", user))
	}
}
