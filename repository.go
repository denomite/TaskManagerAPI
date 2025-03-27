/*
Database interaction (CRUD operations)
  - SetupDatabase initialize the database connection
    AutoMigrate: automatically create or updates the db schema for the Task struct
  - CreateTask insert a new task into database
  - GetAllTask retrieves all task from the database
  - GetTaskByID retrieves a single task by ID
  - UpdateTask updates an existing task
  - DeleteTask removes a task by ID
*/
package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Task{})
	return db
}
func CreateTask(db *gorm.DB, task *Task) *Task {
	db.Create(task)
	return task
}

func GetAllTasks(db *gorm.DB) []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func GetTaskByID(db *gorm.DB, id uint) *Task {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return nil
	}
	return &task
}

func UpdateTask(db *gorm.DB, task *Task) *Task {
	db.Save(task)
	return task
}

func DeleteTask(db *gorm.DB, id uint) {
	db.Delete(&Task{}, id)
}
