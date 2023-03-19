package routers

import (
	"github.com/Gvzum/dias-store.git/api/controllers/auth"
	"github.com/Gvzum/dias-store.git/api/controllers/category"
	"github.com/Gvzum/dias-store.git/api/controllers/product"
	"github.com/Gvzum/dias-store.git/api/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Category handlers
	categoryController := new(category.Controller)
	categoryRoutes := router.Group(
		"/category",
		middlewares.AuthenticationMiddleware(),
		middlewares.ProtectionMiddleware(),
	)
	{
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.GET("/", categoryController.ListCategory)
		categoryRoutes.GET("/:id", categoryController.DetailedCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
	}

	// Product handlers
	productController := new(product.Controller)
	productRoutes := router.Group(
		"/product",
		middlewares.AuthenticationMiddleware(),
		middlewares.ProtectionMiddleware(),
	)
	{
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/", productController.ListProduct)
	}

	// Auth handlers
	authController := new(auth.Controller)
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/sign-up", authController.SignUpHandler)
		authRoutes.POST("/sign-in", authController.SignInHandler)
	}

	return router
}
