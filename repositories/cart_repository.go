package repositories

import (
	"github.com/str122-xyz/gin-firebase-backend/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) GetCartByUserID(userID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.DB.Preload("Items.Product").Where("user_id = ?", userID).FirstOrCreate(&cart, models.Cart{UserID: userID}).Error
	return &cart, err
}

func (r *CartRepository) AddItemToCart(cartID uint, productID uint, quantity int, price float64) error {
	var item models.CartItem
	err := r.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error

	if err == gorm.ErrRecordNotFound {
		newItem := models.CartItem{
			CartID:    cartID,
			ProductID: productID,
			Quantity:  quantity,
			Subtotal:  float64(quantity) * price,
		}
		return r.DB.Create(&newItem).Error
	}

	item.Quantity += quantity
	item.Subtotal = float64(item.Quantity) * price
	return r.DB.Save(&item).Error
}

func (r *CartRepository) UpdateItemQuantity(itemID uint, quantity int) error {
	var item models.CartItem
	if err := r.DB.Preload("Product").First(&item, itemID).Error; err != nil {
		return err
	}
	item.Quantity = quantity
	item.Subtotal = float64(quantity) * item.Product.Price
	return r.DB.Save(&item).Error
}

func (r *CartRepository) DeleteItem(itemID uint) error {
	return r.DB.Delete(&models.CartItem{}, itemID).Error
}

func (r *CartRepository) ClearCart(cartID uint) error {
	return r.DB.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}