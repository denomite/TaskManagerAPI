package controllers

import (
	"TaskManagerAPI/middleware"
	"TaskManagerAPI/models"
	"TaskManagerAPI/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTaskRouter(db *gorm.DB, r *gin.Engine) {
	taskGroup := r.Group("/tasks")

	// Protect all routes in this group with AuthMiddleware
	taskGroup.Use(middleware.AuthMiddleware("user", "admin"))

	// Create task (user can create task)
	taskGroup.POST("/", func(c *gin.Context) {
		var task models.Task
		// Bind the JSON body to the Task struct
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Extract the user ID from the context (from the JWT)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		// Create the task using the repository function
		createdTask, err := repository.CreateTask(db, &task, userID.(uint))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
			return
		}

		c.JSON(http.StatusCreated, createdTask)
	})

	taskGroup.GET("/", func(c *gin.Context) {
		role, _ := c.Get("role")

		var tasks []models.Task
		var err error
		if role == "admin" {
			// Admin can view all tasks
			tasks, err = repository.GetAllTasks(db)
		} else {
			userID, _ := c.Get("userID")
			tasks, err = repository.GetTasksByUserID(db, userID.(uint))
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})

	// Get a single task(admin can view all tasks, user only their tasks)
	taskGroup.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		// Convert the id from string to uint
		taskID, err := strconv.Atoi(id) // Converts string to int

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
			return
		}

		// Now pass TaskID(as uint) to GetTaskByID
		task, err := repository.GetTaskByID(db, uint(taskID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
			return
		}

		// Check if user is allowed to access this task
		role, _ := c.Get("role")
		if role == "user" {
			userID, _ := c.Get("userID")
			if task.UserID != userID.(uint) {
				c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to view this task"})
				return
			}
		}

		c.JSON(http.StatusOK, task)
	})

	taskGroup.PUT("/:id", func(c *gin.Context) {
		// Get the task ID from the URL
		id := c.Param("id")
		taskID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		// Get the user ID and role from the context
		userID, _ := c.Get("userID")
		role, _ := c.Get("role")

		// Retrieve the task from the database
		existingTask, err := repository.GetTaskByID(db, uint(taskID))
		if err != nil || existingTask == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		// Check if the user is an admin or the owner of the task
		if role != "admin" && existingTask.UserID != userID.(uint) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this task"})
			return
		}

		// Bind the updated task data
		var updatedTask *models.Task
		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Update task fields
		existingTask.Title = updatedTask.Title
		existingTask.Description = updatedTask.Description

		// Save the updated task to the database
		updatedTask, err = repository.UpdateTask(db, existingTask)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
			return
		}

		c.JSON(http.StatusOK, updatedTask)
	})

	taskGroup.DELETE("/:id", func(c *gin.Context) {
		// Get the task ID from the URL
		id := c.Param("id")
		taskID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		// Get the user ID and role from the context
		userID, _ := c.Get("userID")
		role, _ := c.Get("role")

		// Retrieve the task from the database
		existingTask, err := repository.GetTaskByID(db, uint(taskID))
		if err != nil || existingTask == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		// Check if the user is an admin or the owner of the task
		if role != "admin" && existingTask.UserID != userID.(uint) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this task"})
			return
		}

		// Delete the task
		err = repository.DeleteTask(db, uint(taskID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	})

}
