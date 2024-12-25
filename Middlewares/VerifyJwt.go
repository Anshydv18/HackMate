package middlewares

import (
	utils "NotesBuddy/Utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token_str, err := c.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, errors.New("unauthorized person"))
			c.Abort()
			return
		}

		tokenJWT, err := jwt.Parse(token_str, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return utils.JWT_KEY, nil
		})

		if err != nil || !tokenJWT.Valid {
			c.JSON(http.StatusUnauthorized, errors.New("invalid token"))
			c.Abort()
			return
		}

		if claims, ok := tokenJWT.Claims.(jwt.MapClaims); ok {
			c.Set("phone", claims["phone"])
		}

		c.Next()
	}
}
