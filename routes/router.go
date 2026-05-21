package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/str122-xyz/gin-firebase-backend/handlers"
	"github.com/str122-xyz/gin-firebase-backend/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Init handlers
	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()

	// API v1 group
	v1 := r.Group("/v1")

	// Health check (tidak perlu auth)
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "gin-firebase-backend"})
	})

	// Auth routes (public)
	auth := v1.Group("/auth")
	auth.POST("/verify-token", authHandler.VerifyToken)

	// Products
	products := v1.Group("/products")
    products.GET("", productHandler.GetAll)
    products.GET("/:id", productHandler.GetByID)

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	
	// Create, Update, Delete hanya untuk admin
	adminProducts := products.Group("")
	adminProducts.Use(middleware.AdminOnly())
	
	adminProducts.POST("", productHandler.Create)
	adminProducts.PUT("/:id", productHandler.Update)
	adminProducts.DELETE("/:id", productHandler.Delete)

	return r
}