/*
Basic CRUD operations for managins tasks
SetupRouter initializes the Gin router and API routes
- API routes to create task, get all tasks, get a task by ID, update task and delete a task
*/
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Create a new task
	r.POST("/tasks", func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
			return
		}

		c.JSON(http.StatusOK, task)
	})

	r.GET("/tasks", func(c *gin.Context) {
		var tasks []Task
		if err := db.Find(&tasks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		var task Task
		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
			return
		}

		c.JSON(http.StatusOK, task)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingTask Task
		if err := db.First(&existingTask, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
			return
		}

		existingTask.Title = task.Title
		existingTask.Description = task.Description
		existingTask.IsCompleted = task.IsCompleted

		if err := db.Save(&existingTask).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
			return
		}

		c.JSON(http.StatusOK, existingTask)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		if err := db.Delete(&Task{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	})

	return r
}
