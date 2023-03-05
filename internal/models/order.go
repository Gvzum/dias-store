package models

import "gorm.io/gorm"

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
