package repository

import (
	"TaskManagerAPI/models"

	"gorm.io/gorm"
)

// CreateUser creates a new user in the database with hashed password and default role
func CreateUser(db *gorm.DB, user *models.User) error {

	// Set default role if not provided
	if user.Role == "" {
		user.Role = "user"
	}

	// Create user in DB
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByUsername retrieves a user from the database by username
func GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
