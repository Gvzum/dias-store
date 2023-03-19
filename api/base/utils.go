package base

import (
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	db := database.GetDB()
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(id string) (*models.User, error) {
	user := models.User{}
	db := database.GetDB()
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
