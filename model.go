/*
Task model - Structure of Task object to store in SQLite database
gorm.model - Predefined struct, to include fields like ID, CreatedAt, UpdatedAt, DeletedAt
Fields - Title of the task, description of the task and boolean to check if task is completed or not
Custom JSON key names added to fileds
*/
package main

import (
	"gorm.io/gorm"
)

type Task struct{
	gorm.Model
		Title string `json:"title"`
		Description string `json:"description"`
		IsCompleted bool `json:"is_completed"`
	}