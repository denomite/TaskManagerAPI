/*
Auth middleware
Creating middleware to check if a user is aut.
Checks if the request has a valid JWT token and extract user_id from the token.
*/
package middleware

import (
	"TaskManagerAPI/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		fmt.Println("🔵 Received Authorization Header:", tokenString) // Log received token

		if tokenString == "" {
			fmt.Println("🔴 Authorization header is missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		fmt.Println("🟢 Extracted Token:", tokenString) // Log extracted token

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			fmt.Println("🔴 Token validation failed:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		fmt.Println("🟢 Token is valid. Extracted UserID:", claims.UserID)
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
