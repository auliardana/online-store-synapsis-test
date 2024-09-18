package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID;references:ID" json:"-"` // Foreign key ke table User

	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID" json:"-"` // Foreign key ke table Product

	PaymentStatus string `gorm:"not null"` // Menyimpan status pembayaran (paid, unpaid)
	TotalPrice    int    `gorm:"not null"`
	Quantity      int    `gorm:"not null"`
}

// type OrderItem struct {
// 	gorm.Model
// 	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

// 	OrderID uuid.UUID `gorm:"type:uuid;not null"`
// 	Order   Order     `gorm:"foreignKey:OrderID;references:ID"` // Foreign key ke table Order

// 	ProductID uuid.UUID `gorm:"type:uuid;not null"`
// 	Product   Product   `gorm:"foreignKey:ProductID;references:ID"` // Foreign key ke table Product

// 	Quantity int `gorm:"not null"`
// 	Price    int `gorm:"not null"`
// }





