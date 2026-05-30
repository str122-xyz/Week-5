package services

import (
	"errors"
	"strconv"

	"github.com/str122-xyz/gin-firebase-backend/models"
	"github.com/str122-xyz/gin-firebase-backend/repositories"
)

type CartService struct {
	CartRepo *repositories.CartRepository
	ProdRepo *repositories.ProductRepository
}

func NewCartService(cartRepo *repositories.CartRepository, prodRepo *repositories.ProductRepository) *CartService {
	return &CartService{
		CartRepo: cartRepo,
		ProdRepo: prodRepo,
	}
}

func (s *CartService) GetCart(userID string) (*models.Cart, float64, int, error) {
	cart, err := s.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, 0, 0, err
	}

	var total float64 = 0
	var itemCount int = 0
	for _, item := range cart.Items {
		total += item.Subtotal
		itemCount += item.Quantity
	}

	return cart, total, itemCount, nil
}

func (s *CartService) AddToCart(userID string, productID uint, quantity int) error {
	// Cek produk valid atau tidak
	product, err := s.ProdRepo.GetProductByID(strconv.Itoa(int(productID)))
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}

	// Dapatkan cart user
	cart, _ := s.CartRepo.GetCartByUserID(userID)

	// Tambahkan ke keranjang
	return s.CartRepo.AddItemToCart(cart.ID, productID, quantity, product.Price)
}

func (s *CartService) UpdateCartItem(itemID uint, quantity int) error {
	return s.CartRepo.UpdateItemQuantity(itemID, quantity)
}

func (s *CartService) RemoveCartItem(itemID uint) error {
	return s.CartRepo.DeleteItem(itemID)
}

func (s *CartService) ClearCart(userID string) error {
	cart, err := s.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	return s.CartRepo.ClearCart(cart.ID)
}