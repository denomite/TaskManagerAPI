// routes/auth_routes.go
package routes

import (
	"TaskManagerAPI/models"
	"TaskManagerAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Login route to authenticate and get a JWT token
	r.POST("/login", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Retrieve the user from the database (replace this with your DB logic)
		var storedUser models.User

		// Check if the password is correct (you can hash and compare)
		if storedUser.Password != user.Password { // In a real app, never store plain-text passwords
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Generate JWT token
		token, err := utils.GenerateJWT(storedUser.ID, storedUser.Role) // Assuming Email is a field in models.User
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	return r
}
