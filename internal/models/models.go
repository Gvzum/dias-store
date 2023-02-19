package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique_index"`
	Password string
}

type Category struct {
	gorm.Model
	Name string
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
	ImageURL    string
	CategoryID  uint
	Category    Category
}

type Order struct {
	gorm.Model
	UserID     uint
	User       User
	ProductID  uint
	Product    Product
	Quantity   int
	TotalPrice float64
	Status     string
}
