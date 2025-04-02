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

// AuthMiddleware checks if the user is auhtneticated
func AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		fmt.Println("🔵 Received Authorization Header:", tokenString) // Log received token

		if tokenString == "" {
			fmt.Println("🔴 Authorization header is missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix if present
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		fmt.Println("🟢 Extracted Token:", tokenString)

		// Validate the JWT token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			fmt.Println("🔴 Token validation failed:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		fmt.Println("🟢 Token is valid. Extracted UserID:", claims.UserID)

		// Set userID and role in context directly from claims
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		if len(requiredRoles) > 0 {
			allowed := false
			for _, r := range requiredRoles {
				if claims.Role == r {
					allowed = true
					break
				}
			}
			if !allowed {
				c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
				c.Abort()
				return
			}

			c.Next()
		}
	}
}
