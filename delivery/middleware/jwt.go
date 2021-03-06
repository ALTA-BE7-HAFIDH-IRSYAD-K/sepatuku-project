package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const secret_jwt = "SECRET"

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(secret_jwt),
	})
}

func CreateToken(id int, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret_jwt))
}

func ReadTokenId(c echo.Context) (int, error) {
	tokenId := c.Get("user").(*jwt.Token)
	if tokenId.Valid {
		claims := tokenId.Claims.(jwt.MapClaims)
		id := int(claims["id"].(float64))
		return id, nil
	}
	return 0, fmt.Errorf("Bad Request")
}
