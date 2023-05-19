package order

import (
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
)

func createOrder(orderData CreateOrderSchema, user *models.User) (*models.Order, error) {
	db := database.GetDB()
	order := models.Order{
		UserID: user.ID,
	}

	if err := db.Create(&order).Error; err != nil {
		return nil, err
	}

	items := make([]models.OrderItem, len(orderData.OrderItems))
	for i, orderItem := range orderData.OrderItems {
		items[i] = models.OrderItem{
			OrderID:   order.ID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
		}
	}

	if err := db.Create(&items).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
