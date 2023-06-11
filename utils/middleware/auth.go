package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Code    int
	Status  string
	Message string
}

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	var jwtKeyByte = []byte(secretKey)
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			result := Auth{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "You are not logged in. Please log in or register first",
			}
			ctx.JSON(http.StatusUnauthorized, result)
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return jwtKeyByte, nil
		})

		if !token.Valid || err != nil {
			result := Auth{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "you are not logged in, please log in or register first",
			}
			ctx.JSON(http.StatusUnauthorized, result)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
