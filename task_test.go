// go test to run tests and go test -v for more verbose output
package main

import (
	"TaskManagerAPI/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	UserID      uint   `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
}

func TestCreateTask(t *testing.T) {
	// dsn := "host=localhost user=testuser password=testpassword dbname=testdb sslmode=disable"
	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}

	// Ensure the table exists
	db.AutoMigrate(&Task{})

	// Test creating a task
	task := Task{Title: "Test Task", Description: "This is a test task"}
	result := db.Create(&task)

	// Ensure the task was created
	assert.NoError(t, result.Error)
	assert.NotZero(t, task.UserID)
	assert.Equal(t, "Test Task", task.Title)
}
