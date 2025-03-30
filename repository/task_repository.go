/*
Database interaction (CRUD operations)
  - SetupDatabase initialize the PostgreSql database connection
    AutoMigrate: automatically create or updates the db schema for the Task struct
  - Updated to Config.go(handles loading env) and .env (to store informations)
  - CreateTask insert a new task into database
  - GetAllTask retrieves all task from the database
  - GetTaskByID retrieves a single task by ID
  - UpdateTask updates an existing task
  - DeleteTask removes a task by ID
*/
package repository

import (
	"log"

	"TaskManagerAPI/config"
	"TaskManagerAPI/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task = models.Task

func SetupDatabase() *gorm.DB {
	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Task{})
	return db
}

func CreateTask(db *gorm.DB, task *models.Task, userID uint) (*models.Task, error) {
	// Set the UserID field in the task
	task.UserID = userID

	// Create the task
	if err := db.Create(task).Error; err != nil {
		return nil, err
	}

	// Preload the associated User when returning the task (to fill the 'User' field)
	var createdTask models.Task
	if err := db.Preload("User").First(&createdTask, task.ID).Error; err != nil {
		return nil, err
	}

	return &createdTask, nil
}
func GetAllTasks(db *gorm.DB) ([]models.Task, error) {
	var tasks []models.Task
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskByID(db *gorm.DB, id uint) (*Task, error) {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func UpdateTask(db *gorm.DB, task *Task) (*Task, error) {
	// We don't want to update the ID or CreatedAt field.
	// We only update the fields that are passed in.
	err := db.Model(task).Where("id = ?", task.ID).Updates(Task{
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
	}).Error

	if err != nil {
		return nil, err
	}

	return task, nil
}

func DeleteTask(db *gorm.DB, id uint) error {
	if err := db.Delete(&Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
