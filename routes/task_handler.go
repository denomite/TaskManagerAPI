/*
Basic CRUD operations for managins tasks
SetupRouter initializes the Gin router and API routes
  - API routes to create task, get all tasks, get a task by ID, update task and delete a task.
  - Check for errors in every database call, which return appropriate HTTP status code, to help
    client(frontend, API users).
    Https status code
    200 OK						Request was successful.
    201 Created 				A resource was successfully created.
    400 Bad Request				The request has invalid input(e.g., invalid JSON, invalid ID).
    404 Not found				The requested resource(task) does not exist.
    500 Internal server Error	Something went wrong on the server(e.g., database error).
*/
package routes

import (
	"TaskManagerAPI/middleware"
	"TaskManagerAPI/models"
	"TaskManagerAPI/repository"
	"TaskManagerAPI/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Task = models.Task

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Protected task routes (require authentication)
	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(middleware.AuthMiddleware()) // 🔒 Protects all task routes
	{
		// Task creation(Admin can create any task, User can only create their own tasks)
		taskRoutes.POST("/", func(c *gin.Context) {
			var task models.Task
			if err := c.ShouldBindJSON(&task); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
				return
			}

			// Extract user ID from JWT token
			userID, err := utils.GetUserIDFromContext(c)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			// Set the UserID to associate the task with the logged-in user
			task.UserID = userID

			// if the user is not a admin, ensure they can only create their own tasks
			role, _ := c.Get("role")
			if role != "admin" {
				task.UserID = userID
			}

			// Create the task
			createdTask, err := repository.CreateTask(db, &task, userID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
				return
			}

			c.JSON(http.StatusCreated, createdTask)
		})

		// Admin can view all tasks, Users can only see their tasks
		taskRoutes.GET("/", func(c *gin.Context) {
			role, _ := c.Get("role")
			var tasks []models.Task
			var err error

			if role == "admin" {
				// Admin can view all tasks
				tasks, err = repository.GetAllTasks(db)
			} else {
				// User can only view their tasks
				userID, _ := c.Get("userID")
				tasks, err = repository.GetTasksByUserID(db, userID.(uint))
			}

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
				return
			}
			c.JSON(http.StatusOK, tasks)
		})
	}

	// Task update route for Admin and User
	r.PUT("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTask models.Task
		taskID, _ := strconv.Atoi(id)

		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Fetch the task from DB
		existingTask, err := repository.GetTaskByID(db, uint(taskID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
			return
		}

		// Check if the user can update this task
		userID, _ := c.Get("userID")
		role, _ := c.Get("role")

		if role != "admin" && existingTask.UserID != userID.(uint) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this task"})
			return
		}

		// Update the task fields
		existingTask.Title = updatedTask.Title
		existingTask.Description = updatedTask.Description
		existingTask.Done = updatedTask.Done

		if _, err := repository.UpdateTask(db, existingTask); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
			return
		}

		c.JSON(http.StatusOK, existingTask)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		taskID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		// Fetch the existing task from DB
		existingTask, err := repository.GetTaskByID(db, uint(taskID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
			return
		}

		if existingTask == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		// Get user ID from context
		userID, _ := c.Get("userID")
		role, _ := c.Get("role")

		// Check if the user is allowed to delete the task
		if role != "admin" && existingTask.UserID != userID.(uint) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this task"})
			return
		}

		err = repository.DeleteTask(db, uint(taskID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	})

	return r
}
