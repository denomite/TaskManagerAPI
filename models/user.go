/*
gorm.model - Predefined struct, to include fields like ID, CreatedAt, UpdatedAt, DeletedAt.
Custom JSON key names added to fields.
*/

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"not null;default:'user'" json:"role"`
	Tasks    []Task `gorm:"foreignKey:UserID"`
}
