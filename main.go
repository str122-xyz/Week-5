package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/str122-xyz/gin-firebase-backend/config"
	"github.com/str122-xyz/gin-firebase-backend/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variable sistem")
	}

	config.InitFirebase()

	config.InitDatabase()

	router := routes.SetupRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di http://localhost:%s", port)
	log.Printf("Health check: http://localhost:%s/v1/health", port)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}