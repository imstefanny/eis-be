package middlewares

import (
	"eis-be/constants"

	"github.com/golang-jwt/jwt"
)

func CreateToken(email, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["password"] = password

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}
