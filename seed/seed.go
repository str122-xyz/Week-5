package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/str122-xyz/gin-firebase-backend/config"
	"github.com/str122-xyz/gin-firebase-backend/models"
)

func main() {
	godotenv.Load()
	config.InitDatabase()

	products := []models.Product{
		// Kategori Kopi
		{Name: "Kopi Susu Ngopss", Price: 25000, Category: "kopi", Stock: 50, Description: "Signature espresso dengan susu segar dan gula aren asli.", ImageURL: "https://images.unsplash.com/photo-1557006021-b85faa2bc5e2?q=80&w=500&auto=format&fit=crop"},
		{Name: "Caffe Latte", Price: 28000, Category: "kopi", Stock: 50, Description: "Kombinasi sempurna espresso dan steamed milk yang lembut.", ImageURL: "https://images.unsplash.com/photo-1572442388796-11668a67e53d?q=80&w=500&auto=format&fit=crop"},
		{Name: "Americano", Price: 20000, Category: "kopi", Stock: 50, Description: "Definisi dari pahitnya kehidupan.", ImageURL: "https://images.unsplash.com/photo-1551030173-122aabc4489c?q=80&w=500&auto=format&fit=crop"},
		
		// Kategori Makanan
		{Name: "Butter Croissant", Price: 22000, Category: "makanan", Stock: 30, Description: "Pastry Prancis klasik yang renyah di luar dan lembut di dalam.", ImageURL: "https://images.unsplash.com/photo-1681218079567-35aef7c8e7e4?q=80&w=1074&auto=format&fit=crop"},
		{Name: "Choco Brownie", Price: 18000, Category: "makanan", Stock: 30, Description: "Fudgy brownie padat dengan potongan dark chocolate.", ImageURL: "https://images.unsplash.com/photo-1606313564200-e75d5e30476c?q=80&w=500&auto=format&fit=crop"},
		{Name: "Pisang Goreng", Price: 15000, Category: "makanan", Stock: 30, Description: "Pisang goreng enak muanis.", ImageURL: "https://images.unsplash.com/photo-1762941904142-9d91ca413e66?q=80&w=735&auto=format&fit=crop"},
	}

	for _, p := range products {
		config.DB.Create(&p)
	}

	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}