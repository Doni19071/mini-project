package loginJWT

import (
	"github.com/golang-jwt/jwt"
)

const SECRET_JWT = "sjjwhke,dsw"

func CreateJwtToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_JWT))
}
