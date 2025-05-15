package helpers

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetTokenClaims(c echo.Context) (jwt.MapClaims, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}
	return claims, nil
}
