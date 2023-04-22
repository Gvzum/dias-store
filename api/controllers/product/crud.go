package product

import (
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
)

func (c Controller) getProductByID(id string) (*models.Product, error) {
	db := database.GetDB()

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
