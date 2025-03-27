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

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	return db
}

func CreateTask(db *gorm.DB, task *Task) *Task {
	db.Create(task)
	return task
}

func GetAllTask(db *gorm.DB) []Task {
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
