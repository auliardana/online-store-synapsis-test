package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name string    `gorm:"unique;not null"`
}

type ProductCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
