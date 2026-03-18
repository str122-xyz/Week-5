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
		{Name: "Nasi Goreng Spesial", Price: 25000, Category: "Makanan", Stock: 50, Description: "Nasi goreng dengan telur dan ayam", ImageURL: "https://picsum.photos/400"},
		{Name: "Sate Padang", Price: 20000, Category: "Makanan", Stock: 100, Description: "Sate padang wuenak cik", ImageURL: "https://picsum.photos/401"},
		{Name: "Mie Ayam", Price: 15000, Category: "Makanan", Stock: 30, Description: "Si my", ImageURL: "https://picsum.photos/402"},
		{Name: "Kopi Hitam", Price: 5000, Category: "Minuman", Stock: 20, Description: "Kopi hitam pahit", ImageURL: "https://picsum.photos/403"},
		{Name: "Es teajus gulawbatu", Price: 2000, Category: "Minuman", Stock: 111, Description: "Es teh manis yang di cekek", ImageURL: "https://picsum.photos/404"},
		{Name: "Air Mineral", Price: 4000, Category: "Minuman", Stock: 67, Description: "Air bening, air putih cakep", ImageURL: "https://picsum.photos/405"},
	}

	for _, p := range products {
		config.DB.Create(&p)
	}

	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}