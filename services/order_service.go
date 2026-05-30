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

	// mengambil riwayat pesanan berdasarkan ID User
func (s *OrderService) GetMyOrders(userID string) ([]models.Order, error) {
	var orders []models.Order
	
	// panggil config.DB untuk mengambil data. 
	// preload("Items.Product") fungsinya untuk detail kopi yang dibeli ikut ketarik.
	// order("created_at desc") untuk pesanan paling baru muncul di paling atas.
	err := config.DB.Preload("Items.Product").Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error
	
	return orders, err
}

	// Mengambil detail satu pesanan berdasarkan Order ID dan User ID
func (s *OrderService) GetOrderByID(orderID string, userID string) (*models.Order, error) {
	var order models.Order
	
	err := config.DB.Preload("Items.Product").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error
		
	if err != nil {
		return nil, errors.New("pesanan tidak ditemukan")
	}
	
	return &order, nil
}