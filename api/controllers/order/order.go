package order

import (
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status int

type Controller struct{}

func (c Controller) CreateOrder(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)

	var validatedOrder CreateOrderSchema
	if err := ctx.ShouldBindJSON(&validatedOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if _, err := createOrder(validatedOrder, user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order successfully created",
	})
}
