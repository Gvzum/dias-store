package product

import (
	"errors"
	"fmt"
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
	"strings"
)

// Product crud

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

func deleteProduct(productID uint, user *models.User) error {
	productSchema, err := getProductByID(productID)
	if err != nil {
		return err
	}
	if productSchema.UserID != user.ID && !user.IsSuperUser {
		return errors.New("don't have permission to delete")
	}
	db := database.GetDB()
	var product models.Product
	if err := db.Delete(&product, productSchema.ID).Error; err != nil {
		return err
	}

	return nil
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

func getListOfProduct(ctx *gin.Context) ([]ListProductSchema, error) {
	db := database.GetDB()

	result := db.Table("products").
		Select("products.*, categories.name as category_name, COALESCE(SUM(product_rates.rate), 0)/NULLIF(COUNT(product_rates.id), 0) as average_rating").
		Joins("LEFT JOIN categories ON categories.id = products.category_id").
		Joins("LEFT JOIN product_rates ON products.id = product_rates.product_id").
		Where("products.deleted_at IS NULL").
		Group("products.id, categories.id")

	if searchName := ctx.Query("name"); searchName != "" {
		searchName = strings.ToLower(searchName)
		result.Where("LOWER(products.name) LIKE ?", "%"+searchName+"%")
	}
	if minPrice, err := strconv.ParseFloat(ctx.Query("min_price"), 64); err == nil {
		result.Where("products.price >= ?", minPrice)
	}

	if maxPrice, err := strconv.ParseFloat(ctx.Query("max_price"), 64); err == nil {
		result.Where("products.price <= ?", maxPrice)
	}

	if minRate, err := strconv.ParseFloat(ctx.Query("min_rate"), 64); err == nil {
		result.Having("COALESCE(SUM(product_rates.rate), 0)/NULLIF(COUNT(product_rates.id), 0) >= ?", minRate)
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

// Product Rate crud

func getCalculatedRate(db *gorm.DB, product *DetailedProductSchema) float64 {
	var sum float64
	var count int64
	result := db.Table("product_rates").
		Select("SUM(rate) as sum, COUNT(*) as count").
		Where("product_id = ? AND deleted_at IS NULL", product.ID).Row()
	err := result.Scan(&sum, &count)
	if err != nil {
		fmt.Println(err.Error())
	}
	if count != 0 {
		return math.Round((sum/float64(count))*100) / 100
	}
	return 0.0
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
	if productRate.UserID != user.ID {
		if !user.IsSuperUser {
			return nil, errors.New("user doesn't have a permission")
		}
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
	if productRate.UserID != user.ID {
		if !user.IsSuperUser {
			return errors.New("doesn't have permission to delete")
		}
	}
	db := database.GetDB()
	if err := db.Delete(&productRate).Error; err != nil {
		return err
	}
	return nil
}

// Comment crud

func getCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func createComment(schema CommentProductSchema, user *models.User) (*models.Comment, error) {
	db := database.GetDB()
	comment := &models.Comment{
		ProductID: schema.ProductID,
		UserID:    user.ID,
		Message:   schema.Message,
	}
	result := db.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}

func updateComment(commentID uint, user *models.User, message string) (*models.Comment, error) {
	comment, err := getCommentByID(commentID)
	if err != nil {
		return nil, err
	}
	if comment.UserID != user.ID {
		if !user.IsSuperUser {
			return nil, errors.New("doesn't have permission to update comment")
		}
	}
	db := database.GetDB()
	comment.Message = message
	if err := db.Save(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func listComments(productID uint) ([]UserCommentSchema, error) {
	db := database.GetDB()
	var comments []UserCommentSchema
	if err := db.
		Model(&models.Comment{}).
		Select("id, message, user_id").
		Where("product_id = ?", productID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func detailedComment(commentID uint) (*BaseCommentSchema, error) {
	db := database.GetDB()
	var comment BaseCommentSchema
	result := db.First(&models.Comment{}, commentID).Scan(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func deleteComment(commentID uint, user *models.User) error {
	comment, err := getCommentByID(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != user.ID {
		if !user.IsSuperUser {
			return errors.New("doesn't have permission to delete")
		}
	}
	db := database.GetDB()
	if err := db.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
