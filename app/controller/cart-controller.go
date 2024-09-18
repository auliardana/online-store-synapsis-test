package controller

import (
	"fmt"
	"net/http"
	"online-store/app/db"
	"online-store/app/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllCart(c *gin.Context) {
	var cart []model.Cart

	if err := db.DBConn.Find(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find cart.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})

}

func CreateCart(c *gin.Context) {
	var cart model.Cart
	var cartRequest model.CartRequest

	if err := c.ShouldBindJSON(&cartRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//parse user uuid
	userUUID, err := uuid.Parse(cartRequest.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user UUID"})
		return
	}

	//parse product uuid
	productUUID, err := uuid.Parse(cartRequest.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product UUID"})
		return
	}

	cart = model.Cart{
		ProductID: productUUID,
		Quantity:  cartRequest.Quantity,
		UserID:    userUUID,
	}

	tx := db.DBConn.Begin()

	if err := tx.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, model.CartResponse{
		Message: "Cart created successfully",
		Data:    cartRequest,
	})
}

func DeleteCartByID(c *gin.Context) {
	var cart model.Cart
	var cartDeleteResponse model.CartDeleteResponse

	cartUUID := c.Param("id")

	cartID, err := uuid.Parse(cartUUID)
	//error handle if role id is not valid
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	tx := db.DBConn.Begin()

	// Load the cart from the database
	if err := tx.Where("id = ?", cartID).First(&cart).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Populate cartRequest with the loaded cart data
	cartDeleteResponse = model.CartDeleteResponse{
		CartID:    cart.ID.String(),
		UserID:    cart.UserID.String(),
		ProductID: cart.ProductID.String(),
		Quantity:  cart.Quantity,
	}

	if err := tx.Where("id = ?", cartID).Delete(&cart).Error; err != nil {
		fmt.Println("lalallalala")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, model.CartResponse{
		Message: "Cart deleted successfully",
		Data:    cartDeleteResponse,
	})
}
