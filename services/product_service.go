package services

import (
	"github.com/str122-xyz/gin-firebase-backend/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{productRepo: repositories.NewProductRepository()}
}