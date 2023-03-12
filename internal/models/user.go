package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;unique_index;not null"`
	Password string
	//Permissions []Permission `gorm:"many2many:user_permissions;" json:"permissions"`
}

type Permission struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_permissions;" json:"users"`
}
