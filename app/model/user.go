package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLoginResponse struct {
	Data  UserResponse `json:"data"`
	Token string       `json:"token"`
}
