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
		c.Set("userID", claims.UserID)

		// Manually map fields from *utils.Claims to mapClaims
		mapClaims := make(map[string]any)
		mapClaims["user_id"] = claims.UserID
		mapClaims["role"] = claims.Role

		// Store userID and role in context
		userID, ok := mapClaims["user_id"].(float64) // Convert to float64
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		role, ok := mapClaims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
			c.Abort()
			return
		}

		// Set userID and role in context
		c.Set("userID", uint(userID))
		c.Set("role", role)

		// Check if user has the required role (if provided)
		if len(requiredRoles) > 0 {
			allowed := false
			for _, r := range requiredRoles {
				if role == r {
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
