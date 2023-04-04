package helpers

import (
	"os"

	"github.com/golang-jwt/jwt"
)

var JWTKEY = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	tokenString, err := token.SignedString(JWTKEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
