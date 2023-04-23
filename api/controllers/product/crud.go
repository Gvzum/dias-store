package product

import (
	"errors"
	"fmt"
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
	"gorm.io/gorm"
	"math"
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

func getProductByID(id uint) (*DetailedProductSchema, error) {
	db := database.GetDB()

	var detailedProduct DetailedProductSchema
	result := db.First(&models.Product{}, id).Scan(&detailedProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	detailedProduct.Rate = getCalculatedRate(db, &detailedProduct)
	return &detailedProduct, nil
}

func getListOfProduct(searchName string) ([]ListProductSchema, error) {
	db := database.GetDB()

	result := db.Table("products").
		Select("products.*, categories.name as category_name").
		Joins("LEFT JOIN categories ON categories.id = products.category_id")

	if searchName != "" {
		searchName = strings.ToLower(searchName)
		result.Where("LOWER(products.name) LIKE ?", "%"+searchName+"%")
	}

	var products []ListProductSchema
	result = result.Scan(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	if products == nil {
		return []ListProductSchema{}, nil
	}
	return products, nil
}

func isProductOwner(rateProductSchema RateProductSchema, user *models.User) bool {
	product, err := getProductByID(rateProductSchema.ProductID)
	if err != nil {
		fmt.Println(err.Error())
	}

	if product.UserID == user.ID {
		return true
	}
	return false
}

func getCalculatedRate(db *gorm.DB, product *DetailedProductSchema) float64 {
	var sum float64
	var count int64
	result := db.Table("product_rates").
		Select("SUM(rate) as sum, COUNT(*) as count").
		Where("product_id = ?", product.ID).Row()
	err := result.Scan(&sum, &count)
	if err != nil {
		fmt.Println(err.Error())
	}
	return math.Round((sum/float64(count))*100) / 100
}

func getRateProduct(productID uint, userID uint) (*models.ProductRate, error) {
	var productRate models.ProductRate
	db := database.GetDB()

	result := db.Where("user_id = ? AND product_id = ?", userID, productID).Last(&productRate)
	if result.Error != nil {
		return nil, result.Error
	}
	return &productRate, nil
}

//func getRateProductByID(productID uint, UserID uint) (*models.ProductRate, error) {
//	var productRate models.ProductRate
//	db := database.GetDB()
//
//	result := db.Where("id = ?", productID).First(&productRate)
//}

func createRateProduct(schema RateProductSchema, user *models.User) (*models.ProductRate, error) {
	isProductOwner := isProductOwner(schema, user)
	if isProductOwner {
		return nil, errors.New("owner is not allowed to rate own products")
	}

	if product, _ := getRateProduct(schema.ProductID, user.ID); product != nil {
		return nil, errors.New("already rated product")
	}

	db := database.GetDB()
	rateProduct := models.ProductRate{
		Rate:      schema.Rate,
		ProductID: schema.ProductID,
		UserID:    user.ID,
	}
	result := db.Create(&rateProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rateProduct, nil
}

func updateRateProduct(schema RateProductSchema, user *models.User) (*models.ProductRate, error) {
	productRate, err := getRateProduct(schema.ProductID, user.ID)
	if err != nil {
		return nil, err
	}
	productRate.Rate = schema.Rate
	db := database.GetDB()
	if err := db.Save(&productRate).Error; err != nil {
		return nil, err
	}
	return productRate, nil
}

func deleteRateProduct(productID uint, user *models.User) error {
	productRate, err := getRateProduct(productID, user.ID)
	if err != nil {
		return err
	}
	if productRate.UserID == user.ID || user.IsSuperUser {
		db := database.GetDB()
		if err := db.Delete(&productRate).Error; err != nil {
			return err
		}
		return nil
	}

	return errors.New("doesn't have permission to delete")
}
