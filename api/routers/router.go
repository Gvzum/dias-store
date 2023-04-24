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
		"api/category",
		middlewares.AuthenticationMiddleware(),
		middlewares.ProtectionMiddleware(),
		middlewares.IsSuperUserMiddleware(),
	)
	{
		categoryRoutes.GET("/:id", categoryController.DetailedCategory)
		categoryRoutes.GET("/", categoryController.ListCategory)
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
	}

	// Product handlers
	productController := new(product.Controller)
	productRoutes := router.Group(
		"api/product",
		middlewares.AuthenticationMiddleware(),
		middlewares.ProtectionMiddleware(),
	)
	{
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/", productController.ListProduct)
		productRoutes.GET("/:id", productController.DetailedProduct)
	}

	// Product Rate handlers
	rateProductController := router.Group(
		"api/product",
		middlewares.AuthenticationMiddleware(),
		middlewares.ProtectionMiddleware(),
	)
	{
		rateProductController.POST("/:id/rate", productController.CreateRateProduct)
		rateProductController.PUT("/:id/rate", productController.UpdateRateProduct)
		rateProductController.DELETE("/:id/rate", productController.DeleteRateProduct)
	}

	// Product Comment handlers
	commentProductController := router.Group(
		"api/product",
		middlewares.AuthenticationMiddleware(),
		middlewares.ProtectionMiddleware(),
	)
	{
		commentProductController.GET("/:id/comment", productController.ListComments)
		commentProductController.POST("/:id/comment", productController.CreateComment)
		commentProductController.PUT("/:id/comment/:comment_id", productController.UpdateComment)
		commentProductController.GET("/:id/comment/:comment_id", productController.DetailedComment)
		commentProductController.DELETE("/:id/comment/:comment_id", productController.DeleteComment)
	}

	// Auth handlers
	authController := new(auth.Controller)
	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/sign-up", authController.SignUpHandler)
		authRoutes.POST("/sign-in", authController.SignInHandler)
	}

	return router
}
