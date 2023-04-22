package product

import (
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strings"
)

type Controller struct{}

func (c Controller) CreateProduct(ctx *gin.Context) {
	var validatedProduct CreateProductSchema
	if err := ctx.ShouldBindJSON(&validatedProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := database.GetDB()
	var category models.Category
	if err := db.First(&category, validatedProduct.CategoryID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	user, _ := ctx.Value("user").(*models.User)

	product := models.Product{
		Name:        validatedProduct.Name,
		Description: validatedProduct.Description,
		Price:       validatedProduct.Price,
		ImageURL:    validatedProduct.ImageURL,
		CategoryID:  validatedProduct.CategoryID,
		UserID:      user.ID,
	}

	if err := db.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product successfully created",
	})

}

func (c Controller) ListProduct(ctx *gin.Context) {
	var products []DetailedProductSchema
	db := database.GetDB()

	result := db.
		Table("products").
		Select("products.*, categories.name as category_name").
		Joins("LEFT JOIN categories ON categories.id = products.category_id")

	if searchName := ctx.Query("name"); searchName != "" {
		searchName = strings.ToLower(searchName)
		result.Where("LOWER(products.name) LIKE ?", "%"+searchName+"%")
	}

	result = result.Scan(&products)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch products",
		})
		return
	}

	if products == nil {
		products = []DetailedProductSchema{}
	}

	ctx.JSON(http.StatusOK, products)
}

func (c Controller) DetailedProduct(ctx *gin.Context) {
	product, err := c.getProductByID(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch product",
		})
		return
	}

	var detailedProduct DetailedProductSchema
	if err := mapstructure.Decode(product, &detailedProduct); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode product",
		})
		return
	}

	ctx.JSON(http.StatusOK, detailedProduct)
}
