package user

import (
	"net/http"
	_middlewares "sepatuku-project/delivery/middleware"
	"sepatuku-project/delivery/response"
	_entities "sepatuku-project/entity/user"
	_userService "sepatuku-project/service/user"
	"strconv"

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
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		users, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("data not exist"))
		}
		if idToken != id {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get data", users))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		_, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		if idToken != id {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
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
		users, err := uh.userService.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed create data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Succes create data", users))
	}
}

func (uh *UserHandler) UpdatedUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		users, rows, err := uh.userService.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Data not exist"))
		}
		if idToken != id {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Bad Request"))
		}
		c.Bind(&users)
		users, err = uh.userService.UpdatedUser(users, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("Failed edit data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Success edit data", users))
	}
}
