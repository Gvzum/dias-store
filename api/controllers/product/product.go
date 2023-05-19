package product

import (
	"fmt"
	"github.com/Gvzum/dias-store.git/api/base"
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct{}

func (c Controller) CreateProduct(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)

	var validatedProduct CreateProductSchema
	if err := ctx.ShouldBindJSON(&validatedProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if _, err := base.GetCategoryByID(validatedProduct.CategoryID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	if _, err := createProduct(validatedProduct, user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product successfully created",
	})

}

func (c Controller) ListProduct(ctx *gin.Context) {
	products, err := getListOfProduct(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch product",
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c Controller) DetailedProduct(ctx *gin.Context) {
	productId := ctx.Param("id")
	if productId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "ID of product doesn't provided",
		})
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	product, err := getProductByID(uint(id))
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Doesn't have this product",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c Controller) DeleteProduct(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)

	productId := ctx.Param("id")
	if productId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "ID of product doesn't provided",
		})
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}
	if err := deleteProduct(uint(id), user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}
