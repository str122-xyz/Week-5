package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/str122-xyz/gin-firebase-backend/services"
)

type CartHandler struct {
	CartService *services.CartService
}

// GET /v1/cart
func (h *CartHandler) GetCart(c *gin.Context) {
	userID := c.MustGet("firebase_uid").(string)

	cart, total, itemCount, err := h.CartService.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil keranjang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"items":      cart.Items,
			"total":      total,
			"item_count": itemCount,
		},
	})
}

// POST /v1/cart
func (h *CartHandler) AddToCart(c *gin.Context) {
	userID := c.MustGet("firebase_uid").(string)

	var req struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format request salah"})
		return
	}

	if err := h.CartService.AddToCart(userID, req.ProductID, req.Quantity); err != nil {
		if err.Error() == "produk tidak ditemukan" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan ke keranjang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil ditambahkan ke keranjang"})
}

// PUT /v1/cart/:id
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Quantity int `json:"quantity"`
	}
	c.ShouldBindJSON(&req)

	h.CartService.UpdateCartItem(uint(itemID), req.Quantity)
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil diupdate"})
}

// DELETE /v1/cart/:id
func (h *CartHandler) RemoveCartItem(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.Param("id"))
	h.CartService.RemoveCartItem(uint(itemID))
	c.JSON(http.StatusOK, gin.H{"message": "Item dihapus"})
}

// DELETE /v1/cart
func (h *CartHandler) ClearCart(c *gin.Context) {
	userID := c.MustGet("firebase_uid").(string)
	h.CartService.ClearCart(userID)
	c.JSON(http.StatusOK, gin.H{"message": "Keranjang dibersihkan"})
}