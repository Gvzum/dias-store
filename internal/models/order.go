package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	User   User
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	Order     Order
	ProductID uint
	Product   Product
	Quantity  int
}
