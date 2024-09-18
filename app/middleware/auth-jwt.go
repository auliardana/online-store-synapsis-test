package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Ambil token dari header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization header is missing",
		})
		c.Abort()
		return
	}

	// Token harus menggunakan skema "Bearer {token}"
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization header format must be Bearer {token}",
		})
		c.Abort()
		return
	}

	tokenString := tokenParts[1]

	// Parse dan validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi metode signing token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Ambil secret key dari environment
		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			return nil, fmt.Errorf("secret key is not set in environment variables")
		}

		return []byte(secretKey), nil
	})

	// Jika token tidak valid atau terjadi error
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		c.Abort()
		return
	}

	// Ambil claims dari token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token claims",
		})
		c.Abort()
		return
	}

	// Cek apakah token sudah expired
	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token has expired",
		})
		c.Abort()
		return
	}

	// Ambil user_id dari claims (diasumsikan user_id adalah string)
	userID, ok := claims["user_id"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user ID format",
		})
		c.Abort()
		return
	}

	// Simpan user_id ke dalam context Gin
	c.Set("user_id", userID)

	// Lanjut ke handler berikutnya
	c.Next()
}
