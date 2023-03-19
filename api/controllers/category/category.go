package category

import (
	"fmt"
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func (c Controller) CreateCategory(ctx *gin.Context) {
	var validatedCategory CreateCategorySchema
	if err := ctx.ShouldBindJSON(&validatedCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := database.GetDB()
	var category models.Category
	db.FirstOrCreate(&category, models.Category{Name: validatedCategory.Name})

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
	})

}

func (c Controller) ListCategory(ctx *gin.Context) {
	var categories []DetailedCategorySchema
	db := database.GetDB()
	if err := db.Model(&models.Category{}).Select("id, name").Scan(&categories).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	fmt.Println(categories)
	ctx.JSON(http.StatusOK, categories)
}

func (c Controller) DetailedCategory(ctx *gin.Context) {
	var category DetailedCategorySchema
	db := database.GetDB()
	if err := db.Model(&models.Category{}).Select("id, name").First(&category, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c Controller) DeleteCategory(ctx *gin.Context) {
	var category models.Category
	db := database.GetDB()
	if err := db.First(&category, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	db.Delete(&category)
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

func (c Controller) UpdateCategory(ctx *gin.Context) {
	var category models.Category
	db := database.GetDB()
	if err := db.First(&category, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&category)
	ctx.JSON(http.StatusOK, category)
}
