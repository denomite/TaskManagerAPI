/*
Entry point of application.
It initializes the database connection, set up the routes and start the server on port 8080
*/
package main

import (
	"TaskManagerAPI/controllers"
	"TaskManagerAPI/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database
	db, err := gorm.Open(sqlite.Open("task_manager.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Migrate the schema (create tables)
	db.AutoMigrate(&models.User{}, &models.Task{})

	// Set up the router
	r := gin.Default()

	// Pass DB to routes
	controllers.SetupAuthRouter(db, r)
	controllers.SetupTaskRouter(db, r)

	r.Run(":8080")
}
