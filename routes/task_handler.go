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
	"TaskManagerAPI/models"
	"TaskManagerAPI/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Task = models.Task

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		createdTask, err := repository.CreateTask(db, &task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
			return
		}

		c.JSON(http.StatusCreated, createdTask)
	})

	r.GET("/tasks", func(c *gin.Context) {
		tasks, err := repository.GetAllTasks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		var task *Task
		id := c.Param("id")

		taskID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		task, err = repository.GetTaskByID(db, uint(taskID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
			return
		}

		if task == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		c.JSON(http.StatusOK, task)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		taskID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		var updatedTask *Task
		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		existingTask, err := repository.GetTaskByID(db, uint(taskID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
			return
		}

		if existingTask == nil {
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

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		taskID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		existingTask, err := repository.GetTaskByID(db, uint(taskID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve task"})
			return
		}

		if existingTask == nil {
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

	return r
}
