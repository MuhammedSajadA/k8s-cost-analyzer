package middleware

import (
	"net/http"
	"strings"
	"fmt"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		// c.Set("user_id", claims["user_id"].(string))
		userID, ok := claims["user_id"]
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "user_id missing in token"})
			return
		}

		c.Set("user_id", fmt.Sprintf("%v", userID))

		c.Next()
	}
}
