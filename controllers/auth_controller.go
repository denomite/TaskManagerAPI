// Handles user authentication (register/login)
package controllers

import (
	"TaskManagerAPI/models"
	"TaskManagerAPI/repository"
	"TaskManagerAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupAuthRouter sets up the routes for user authentication (Register and Login)
func SetupAuthRouter(db *gorm.DB, r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", Register)
		authGroup.POST("/login", Login)
	}
}

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash password before saving to DB
	hashedPassword, _ := utils.HashPassword(input.Password)
	user := models.User{Username: input.Username, Password: hashedPassword}

	// Retrieve the db connection from the context
	db := c.MustGet("db").(*gorm.DB)

	// Call repository to create the user
	if err := repository.CreateUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Retrieve the db connection from the context
	db := c.MustGet("db").(*gorm.DB)

	// Retrieve the user by username
	user, err := repository.GetUserByUsername(db, input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the password is correct
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, _ := utils.GenerateJWT(user.ID, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
