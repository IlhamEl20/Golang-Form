package middleware

import (
	"Si-KP/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// const SecretKey = constants.SECRET_KEY
var SecretKey = []byte(config.Env("SECRET_KEY"))

//  SecretKey :=config.Env("SECRET_KEY")

// untuk generate token ketika berhasil login
func GenerateJwt(issuer string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer,
		// ExpiresAt: time.Now().Add(time.Minute * 3).Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(SecretKey))
}
