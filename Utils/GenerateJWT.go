package utils

import (
	constants "Hackmate/Constants"
	env "Hackmate/Env"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(env.Get(constants.JWT_TOKEN))

func GenerateJWTkey(Phone string) (string, error) {
	claims := jwt.MapClaims{
		"phone":  Phone,
		"expire": time.Now().Add(1 * time.Hour).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(JWT_KEY)
}
