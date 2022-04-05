package auth

import (
	"fmt"
	"net/http"
	"sepatuku-project/delivery/response"
	_authService "sepatuku-project/service/auth"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService _authService.AuthServiceInterface
}

func NewAuthHandler(auth _authService.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: auth,
	}
}

func (ah *AuthHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		type loginData struct {
			Identifier string `json:"identifier"`
			Password   string `json:"password"`
		}
		var login loginData
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("error bind data"))
		}
		token, errorLogin := ah.authService.Login(login.Identifier, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("login succesfully", responseToken))
	}
}
