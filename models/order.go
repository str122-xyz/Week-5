package models

import "time"

type Order struct {
	ID              uint        `json:"id" gorm:"primaryKey"`
	UserID          string      `json:"user_id" gorm:"index"`
	TotalAmount     float64     `json:"total_amount"`
	PaymentMethod   string      `json:"payment_method"`
	ShippingAddress string      `json:"shipping_address"`
	Status          string      `json:"status"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id" gorm:"index"`
	ProductID uint      `json:"product_id"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Subtotal  float64   `json:"subtotal"`
}