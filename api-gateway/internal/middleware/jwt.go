package middleware

import (
	"net/http"
	"strings"

	"api-gateway/pkg/jwt"
	"github.com/gin-gonic/gin"
)

const (
	ContextUserID = "user_id"
	ContextEmail  = "email"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// валидная ли структура токена (не расшифровка)
		if !jwt.Sanitize(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token structure"})
			return
		}
		secretKey := jwt.ProvideSecretKey()
		// расшифровка
		claims, err := jwt.Decode(token, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// Прокидываем в контекст
		c.Set(ContextUserID, claims.UUID)
		c.Set(ContextEmail, claims.Email)

		c.Next()
	}
}
