package services

import (
	"errors"

	"github.com/str122-xyz/gin-firebase-backend/config"
	"github.com/str122-xyz/gin-firebase-backend/models"
)

type OrderService struct {
	CartService *CartService
}

func NewOrderService(cartService *CartService) *OrderService {
	return &OrderService{CartService: cartService}
}

func (s *OrderService) Checkout(userID string, address string, payment string) (*models.Order, error) {
	// 1. Ambil keranjang user saat ini
	cart, total, _, err := s.CartService.GetCart(userID)
	if err != nil || len(cart.Items) == 0 {
		return nil, errors.New("keranjang kosong atau gagal diambil")
	}

	// 2. Buat data order baru
	order := models.Order{
		UserID:          userID,
		TotalAmount:     total,
		PaymentMethod:   payment,
		ShippingAddress: address,
		Status:          "Menunggu Pembayaran",
	}

	tx := config.DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 3. Pindahkan item dari keranjang ke order item
	for _, cartItem := range cart.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Product.Price,
			Subtotal:  cartItem.Subtotal,
		}
		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 4. Bersihkan keranjang user karena sudah di-checkout
	if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &order, nil
}