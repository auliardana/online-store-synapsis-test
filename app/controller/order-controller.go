package controller

import (
	"net/http"
	"online-store/app/model"

	"online-store/app/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllOrder(c *gin.Context) {
	var orders []model.Order

	// Ambil user_id dari JWT (asumsikan sudah ada middleware JWT yang menyimpan user_id di context)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not authorized",
		})
		return
	}

	// Konversi userID menjadi UUID
	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user ID format",
		})
		return
	}

	// Query hanya orders yang sesuai dengan user_id
	if err := db.DBConn.Where("user_id = ?", userUUID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't find orders.",
			"error":   err.Error(),
		})
		return
	}

	// Kembalikan response jika data ditemukan
	c.JSON(http.StatusOK, gin.H{
		"message": "Orders found",
		"orders":  orders,
	})
}
