package routers

import (
	"github.com/Gvzum/dias-store.git/api/controllers/auth"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Auth handlers
	authController := new(auth.Controller)
	router.POST("/sign-up", authController.SignUpHandler)
	router.POST("/sign-in", authController.SignInHandler)

	return router
}
