/*
Task model - Structure of Task object to store.
gorm.model - Predefined struct, to include fields like ID, CreatedAt, UpdatedAt, DeletedAt.
Fields - Title of the task, description of the task and boolean to check if task is completed or not.
Custom JSON key names added to fileds.
Include User model and update the Task model.
*/
package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
	Tasks    []Task `gorm:"foreignKey:UserId"`
}

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
	UserID      uint   `json:"user_id"`
}
