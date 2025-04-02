// go test to run tests and go test -v for more verbose output
package main

import (
	"TaskManagerAPI/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	UserID      uint   `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
}

func TestCreateTask(t *testing.T) {
	// Set up a test database connection
	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}

	// Ensure the table exists
	db.AutoMigrate(&User{}, &Task{})

	// Create a test user
	user := User{Username: "testuser", Password: "testpassword"}
	result := db.Create(&user)
	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	// Now create a task associated with the user
	task := Task{UserID: user.ID, Title: "Test Task", Description: "This is a test task"}
	result = db.Create(&task)

	// Ensure the task was created
	assert.NoError(t, result.Error)
	assert.NotZero(t, task.UserID)                           // Ensure that the task has a valid UserID
	assert.Equal(t, "Test Task", task.Title)                 // Check that the title matches
	assert.Equal(t, "This is a test task", task.Description) // Check description
}
