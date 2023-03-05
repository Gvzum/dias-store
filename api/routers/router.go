package routers

import (
	"github.com/Gvzum/dias-store.git/api/controllers/auth"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//productController := new(product.Controller)
	userController := new(auth.Controller)

	router.POST("/signup", userController.Login)
	router.POST("/login")

	return router
}
