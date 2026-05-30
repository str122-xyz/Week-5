package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/str122-xyz/gin-firebase-backend/config"
	"github.com/str122-xyz/gin-firebase-backend/handlers"
	"github.com/str122-xyz/gin-firebase-backend/middleware"
	"github.com/str122-xyz/gin-firebase-backend/repositories"
	"github.com/str122-xyz/gin-firebase-backend/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	// init handlers & repositories (Product & Auth)
	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()

	// init handler cart
	cartRepo := repositories.NewCartRepository(config.DB)
	prodRepo := repositories.NewProductRepository()
	cartService := services.NewCartService(cartRepo, prodRepo)
	cartHandler := handlers.CartHandler{
		CartService: cartService,
	}
	
	// init order handler
	orderService := services.NewOrderService(cartService) 
	orderHandler := handlers.OrderHandler{
		OrderService: orderService,
	}
	
	// API v1 group
	v1 := r.Group("/v1")

	// Health check (tidak perlu auth)
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "gin-firebase-backend"})
	})

	// Auth routes (public)
	auth := v1.Group("/auth")
	auth.POST("/verify-token", authHandler.VerifyToken)

	// Products routes (public)
	products := v1.Group("/products")
	products.GET("", productHandler.GetAll)
	products.GET("/:id", productHandler.GetByID)

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	
	// route cart
	cart := protected.Group("/cart")
	{
		cart.GET("", cartHandler.GetCart)
		cart.POST("", cartHandler.AddToCart)
		cart.PUT("/:id", cartHandler.UpdateCartItem)
		cart.DELETE("/:id", cartHandler.RemoveCartItem)
		cart.DELETE("", cartHandler.ClearCart)
	}

	// Route Orders
	orders := protected.Group("/orders")
	{
		orders.POST("/checkout", orderHandler.Checkout)
	}

	// Create, Update, Delete hanya untuk admin
	adminProducts := products.Group("")
	adminProducts.Use(middleware.AdminOnly())
	
	adminProducts.POST("", productHandler.Create)
	adminProducts.PUT("/:id", productHandler.Update)
	adminProducts.DELETE("/:id", productHandler.Delete)

	return r
}