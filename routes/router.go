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