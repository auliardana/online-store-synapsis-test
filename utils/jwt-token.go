package utils

// import (
// 	"time"
// 	"github.com/golang-jwt/jwt/v5"
// 	"errors"
// )

// type Claims struct {
// 	UserID uint `json:"user_id"`
// 	jwt.StandardClaims
// }

// var jwtKey = []byte("secret_key") // Kunci rahasia untuk JWT


// // ValidateToken validates the JWT token and returns the claims
// func ValidateToken(tokenString string) (*Claims, error) {
// 	claims := &Claims{}

// 	// Parse token dengan kunci yang sama dan claims yang sesuai
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})

// 	// Cek apakah parsing token menghasilkan error
// 	if err != nil {
// 		if errors.Is(err, jwt.ErrTokenExpired) {
// 			return nil, errors.New("token expired")
// 		}
// 		return nil, err
// 	}

// 	// Cek apakah token valid
// 	if !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}

// 	return claims, nil
// }


// // GenerateAccessToken generates an access token that lasts for 15 minutes.
// func GenerateAccessToken(userID uint) (string, error) {
// 	// Masa berlaku access token
// 	expirationTime := time.Now().Add(15 * time.Minute)

// 	// Membuat claims dengan UserID dan waktu kadaluarsa
// 	claims := &Claims{
// 		UserID: userID,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}

// 	// Membuat JWT token dengan metode signing HMAC SHA256
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Mengembalikan token yang sudah signed
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// // GenerateRefreshToken generates a refresh token that lasts for 7 days.
// func GenerateRefreshToken(userID uint) (string, error) {
// 	// Masa berlaku refresh token
// 	expirationTime := time.Now().Add(7 * 24 * time.Hour)

// 	// Membuat claims dengan UserID dan waktu kadaluarsa
// 	claims := &Claims{
// 		UserID: userID,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}

// 	// Membuat JWT token dengan metode signing HMAC SHA256
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Mengembalikan token yang sudah signed
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
