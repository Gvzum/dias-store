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
	products, err := getListOfProduct(ctx.Query("name"))

	if err != nil {
		fmt.Println(err)
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

// Product Rate Endpoints

func (c Controller) CreateRateProduct(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)
	var rateProductSchema RateProductSchema
	if err := ctx.ShouldBindJSON(&rateProductSchema); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}
	rateProductSchema.ProductID = uint(id)
	if _, err := createRateProduct(rateProductSchema, user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product rated successfully",
	})
}

func (c Controller) UpdateRateProduct(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)
	var rateProductSchema RateProductSchema
	if err := ctx.ShouldBindJSON(&rateProductSchema); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idStr := ctx.Param("id")
	if id, err := strconv.ParseUint(idStr, 10, 0); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	} else {
		rateProductSchema.ProductID = uint(id)
	}

	_, err := updateRateProduct(rateProductSchema, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update rate successfully",
	})
}

func (c Controller) DeleteRateProduct(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)
	idStr := ctx.Param("id")
	productID, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	if err := deleteRateProduct(uint(productID), user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted",
	})
}
