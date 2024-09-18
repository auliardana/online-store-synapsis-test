package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID;references:ID" json:"-"` // Foreign key ke table User

	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID" json:"-"` // Foreign key ke table Product

	Quantity int `gorm:"not null"`
}

type CartDeleteResponse struct {
	CartID   string `json:"cart_id"`
	UserID    string `json:"user_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type CartRequest struct {
	// UserID    string `json:"user_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type CartResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

