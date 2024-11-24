package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goMarket/internal"
)

// JWTMiddleware проверяет наличие токена и его действительность
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Разбор и валидация токена
		token, err := jwt.ParseWithClaims(
			tokenString,
			&jwt.StandardClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return internal.JWTSecretKey, nil
			})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Сохраняем ID пользователя в контексте
		c.Set("userID", claims.Subject)
		c.Next()
	}
}
