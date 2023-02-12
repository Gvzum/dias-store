package models

import "gorm.io/gorm"

type BaseModel struct {
	ID uint `gorm:"primarykey"`
}

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Category struct {
	BaseModel
	Name string
}

type Product struct {
	BaseModel
	Name string
}

type Cart struct {
	BaseModel
}
