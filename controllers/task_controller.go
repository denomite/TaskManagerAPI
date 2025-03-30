package controllers

import (
	"TaskManagerAPI/models"
	"TaskManagerAPI/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTaskRouter(db *gorm.DB, r *gin.Engine) {
	taskGroup := r.Group("/tasks")
	{
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
			tasks, err := repository.GetAllTasks(db)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
				return
			}
			c.JSON(http.StatusOK, tasks)
		})

		taskGroup.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			taskID, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
				return
			}

			task, err := repository.GetTaskByID(db, uint(taskID))
			if err != nil || task == nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
				return
			}

			c.JSON(http.StatusOK, task)
		})

		taskGroup.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			taskID, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
				return
			}

			var updatedTask *models.Task
			if err := c.ShouldBindJSON(&updatedTask); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
				return
			}

			existingTask, err := repository.GetTaskByID(db, uint(taskID))
			if err != nil || existingTask == nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
				return
			}

			updatedTask.ID = existingTask.ID
			updatedTask, err = repository.UpdateTask(db, updatedTask)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
				return
			}
			c.JSON(http.StatusOK, updatedTask)
		})

		taskGroup.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			taskID, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
				return
			}

			existingTask, err := repository.GetTaskByID(db, uint(taskID))
			if err != nil || existingTask == nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
				return
			}

			err = repository.DeleteTask(db, uint(taskID))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
		})
	}
}
