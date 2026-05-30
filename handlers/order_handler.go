package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/str122-xyz/gin-firebase-backend/services"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	userID := c.MustGet("firebase_uid").(string)

	var req struct {
		ShippingAddress string `json:"shipping_address" binding:"required"`
		PaymentMethod   string `json:"payment_method" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Alamat dan metode pembayaran wajib diisi"})
		return
	}

	order, err := h.OrderService.Checkout(userID, req.ShippingAddress, req.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Checkout berhasil",
		"data":    order,
	})
}

	// GET /v1/orders
func (h *OrderHandler) GetOrders(c *gin.Context) {
	// mengambil UID dari token yang sedang login
	userID := c.MustGet("firebase_uid").(string)

	orders, err := h.OrderService.GetMyOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil riwayat pesanan"})
		return
	}

	// mengirim JSON sukses ke Flutter
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil riwayat pesanan",
		"data":    orders,
	})
}

	// GET /v1/orders/:id
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.MustGet("firebase_uid").(string)

	order, err := h.OrderService.GetOrderByID(orderID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil detail pesanan",
		"data":    order,
	})
}