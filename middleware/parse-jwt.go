package middleware

import (
	
	"github.com/golang-jwt/jwt/v4"
)

// untuk check token, dari middleware
func Parsejwt(jwt_token string) (string, error){

	token, err:=jwt.ParseWithClaims(jwt_token, &jwt.StandardClaims{}, func(t *jwt.Token)(interface{}, error){
		return []byte(SecretKey), nil
	})
	
	if err != nil || !token.Valid{
		return "", err
	}

	Claims := token.Claims.(*jwt.StandardClaims)
	return Claims.Issuer, nil

}
