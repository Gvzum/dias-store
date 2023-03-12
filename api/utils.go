package api

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
