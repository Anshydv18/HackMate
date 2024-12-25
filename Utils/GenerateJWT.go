package utils

import (
	constants "NotesBuddy/Constants"
	env "NotesBuddy/Env"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(env.Get(constants.JWT_TOKEN))

func generateJWTkey(Phone string) (string, error) {
	claims := jwt.MapClaims{
		"phone":  Phone,
		"expire": time.Now().Add(1 * time.Hour).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return jwtToken.SignedString(JWT_KEY)
}
