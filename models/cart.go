package models

import (
	"time"
)

type Cart struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    string     `json:"user_id" gorm:"index"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CartID    uint      `json:"cart_id" gorm:"index"`
	ProductID uint      `json:"product_id"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity"`
	Subtotal  float64   `json:"subtotal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}