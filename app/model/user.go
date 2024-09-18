package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FirstName string    `gorm:"unique;not null"`
	LastName  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Phone     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
}

type UserRegisterRequest struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone   string `json:"phone"`
}

type UserRegisterResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

type UserLoginResponse struct {
	Data  UserResponse `json:"data"`
	Token string       `json:"token"`
}
