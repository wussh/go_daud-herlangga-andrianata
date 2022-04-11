package kentang

import (
	"github.com/golang-jwt/jwt"
)

const SecretJwt = "kentang"

func CreateJwtToken(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretJwt))
}
