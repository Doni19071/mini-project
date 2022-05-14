package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_JWT = "sjjwhke,dsw"

func CreateToken(username string, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["authorized"] = true
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_JWT))
}
