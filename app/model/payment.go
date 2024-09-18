package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID;references:ID"` // Foreign key ke table User

	OrderID uuid.UUID `gorm:"type:uuid;not null"`
	Order   Order     `gorm:"foreignKey:OrderID;references:ID"` // Foreign key ke table Order

	PaymentAmount int    `gorm:"not null"`
	PaymentMethod string `gorm:"not null"`
	Status        string
}

type PaymentRequest struct {
	OrderID string `json:"order_id" binding:"required"`
}
