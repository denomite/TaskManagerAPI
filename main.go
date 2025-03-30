/*
Entry point of application.
It initializes the database connection, set up the routes and start the server on port 8080
*/
package main

import (
	"TaskManagerAPI/config"
	"TaskManagerAPI/controllers"
	"TaskManagerAPI/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Middleware to add db to the context
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // Set database connection to the context
		c.Next()
	})

	// Setup authentication and task routers
	controllers.SetupAuthRouter(db, r)
	controllers.SetupTaskRouter(db, r)

	return r
}

func main() {
	// Initialize the database
	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Migrate the schema (create tables)
	db.AutoMigrate(&models.User{}, &models.Task{})

	// Set up and run the server
	r := SetupRouter(db)
	r.Run(":8080")
}
