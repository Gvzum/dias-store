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
	detailedProduct.Rate = getRateProduct(db, &detailedProduct)
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

func getRateProduct(db *gorm.DB, product *DetailedProductSchema) float64 {
	var sum float64
	var count int64
	result := db.Table("product_rates").
		Select("SUM(rate) as sum, COUNT(*) as count").
		Where("product_id = ?", product.ID).Row()
	err := result.Scan(&sum, &count)
	if err != nil {
		fmt.Println(err.Error())
	}
	//return sum / float64(count)
	//return strconv.FormatFloat(10.900, 'f', -1, 64)
	return math.Round((sum/float64(count))*100) / 100
}

func createRateProduct(rateProductSchema RateProductSchema, user *models.User) (*models.ProductRate, error) {
	isProductOwner := isProductOwner(rateProductSchema, user)
	if isProductOwner {
		return nil, errors.New("owner is not allowed to rate own products")
	}
	db := database.GetDB()
	rateProduct := models.ProductRate{
		Rate:      rateProductSchema.Rate,
		ProductID: rateProductSchema.ProductID,
		UserID:    user.ID,
	}
	result := db.Create(&rateProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rateProduct, nil
}
