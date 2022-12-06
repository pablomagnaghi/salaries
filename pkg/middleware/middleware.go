package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"salaries/pkg/auth"
)

type JwtClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func NewAuthMiddleware(authService auth.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := authService.VerifyToken(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			fmt.Println(err)
			context.Abort()
			return
		}
		context.Next()
	}
}
