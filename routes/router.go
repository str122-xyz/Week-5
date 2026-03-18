package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/str122-xyz/gin-firebase-backend/handlers"
	"github.com/str122-xyz/gin-firebase-backend/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS Middleware (izinkan request dari aplikasi Flutter)
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

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

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())

	// Products
	products := protected.Group("/products")
	
	// GET semua user terautentikasi bisa akses
	products.GET("", productHandler.GetAll)
	products.GET("/:id", productHandler.GetByID)

	// Create, Update, Delete hanya untuk admin
	adminProducts := products.Group("")
	adminProducts.Use(middleware.AdminOnly())
	
	adminProducts.POST("", productHandler.Create)
	adminProducts.PUT("/:id", productHandler.Update)
	adminProducts.DELETE("/:id", productHandler.Delete)

	return r
}