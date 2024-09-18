package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	CategoryID uuid.UUID       `gorm:"type:uuid;not null"`
	Category   ProductCategory `gorm:"foreignKey:CategoryID;references:ID" json:"-"` // Foreign key ke table Category

	Name        string `gorm:"unique;not null"`
	ImageURL    string `gorm:"type:varchar(255)"` // URL gambar produk
	Price       int    `gorm:"not null"`
	Stock       int    `gorm:"not null"`
	Description string `gorm:"not null"`
}

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	ImageURL    string `json:"image_url"` // URL gambar produk
	Stock       int    `json:"stock" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category_id" binding:"required"`
}

type ProductResponse struct {
	Message string         `json:"message"`
	Data    ProductRequest `json:"data"`
}


type ProductDeleteResponse struct {
	Message string `json:"message"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
}
