package product

import (
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
	"strings"
)

func createProduct(productData CreateProductSchema, user *models.User) (*models.Product, error) {
	db := database.GetDB()

	product := models.Product{
		Name:        productData.Name,
		Description: productData.Description,
		Price:       productData.Price,
		ImageURL:    productData.ImageURL,
		CategoryID:  productData.CategoryID,
		UserID:      user.ID,
	}
	result := db.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func getProductByID(id string) (*DetailedProductSchema, error) {
	db := database.GetDB()

	var detailedProduct DetailedProductSchema
	result := db.First(&models.Product{}, id).Scan(&detailedProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	return &detailedProduct, nil
}

func getListOfProduct(searchName string) ([]DetailedProductSchema, error) {
	db := database.GetDB()

	result := db.Table("products").
		Select("products.*, categories.name as category_name").
		Joins("LEFT JOIN categories ON categories.id = products.category_id")

	if searchName != "" {
		searchName = strings.ToLower(searchName)
		result.Where("LOWER(products.name) LIKE ?", "%"+searchName+"%")
	}

	var products []DetailedProductSchema
	result = result.Scan(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	if products == nil {
		return []DetailedProductSchema{}, nil
	}
	return products, nil
}
