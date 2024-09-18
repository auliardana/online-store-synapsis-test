package controller

import (
	"net/http"
	"online-store/app/db"
	"online-store/app/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAllCart is a method to get Cart list
// @Summary see a list of products that have been added to the shopping cart
// @Description get a list of all products that have been added to the shopping cart
// @Tags Cart
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer 'Add access token here')
// @Success 200 {object} model.CartResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cart [get]
func GetAllCart(c *gin.Context) {
	// Ambil user_id dari JWT (asumsikan sudah ada middleware JWT yang menyimpan user_id di context)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not authorized",
		})
		return
	}

	// Konversi userID menjadi tipe yang sesuai, misalnya uint atau int
	userIDUint, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user ID",
		})
		return
	}

	// Inisialisasi variabel cart yang akan diisi dengan hasil query
	var cart []model.Cart

	// Query database berdasarkan user_id
	if err := db.DBConn.Where("user_id = ?", userIDUint).Find(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find cart.",
			"error":   err.Error(),
		})
		return
	}

	// Kembalikan hasil query ke client
	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}

// Create is a method to Add product to cart
// @Summary Add product to cart
// @Description Add product to cart
// @Tags Cart
// @Accept  json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer 'Add access token here')
// @Param   cart body model.CartRequest true "cart need product_id and quantity"
// @Success 201 {object} model.CartResponse	"created successfully"
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cart [post]
func CreateCart(c *gin.Context) {
	var cart model.Cart
	var cartRequest model.CartRequest

	if err := c.ShouldBindJSON(&cartRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil user_id dari JWT (asumsikan sudah ada middleware JWT yang menyimpan user_id di context)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not authorized",
		})
		return
	}

	// Konversi userID menjadi tipe yang sesuai, misalnya uint atau int
	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user ID",
		})
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

// Delete is a method to delete a product cart by its id
// @Summary Delete a product cart by id
// @Description Delete a product by id
// @Tags Cart
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer 'Add access token here')
// @Param id path string true "cart ID"
// @Success 200 {object} model.CartResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cart/{id} [delete]
func DeleteCartByID(c *gin.Context) {
	var cart model.Cart
	var cartDeleteResponse model.CartDeleteResponse

	// Ambil UUID dari parameter URL
	cartUUID := c.Param("id")

	// Ambil user_id dari context (disimpan oleh middleware JWT)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not authorized",
		})
		return
	}

	// Konversi cartUUID menjadi UUID
	cartID, err := uuid.Parse(cartUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	// Mulai transaksi database
	tx := db.DBConn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Cari cart berdasarkan id dan user_id
	if err := tx.Where("id = ? AND user_id = ?", cartID, userID).First(&cart).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Siapkan response setelah cart dihapus
	cartDeleteResponse = model.CartDeleteResponse{
		CartID:    cart.ID.String(),
		UserID:    cart.UserID.String(),
		ProductID: cart.ProductID.String(),
		Quantity:  cart.Quantity,
	}

	// Hapus cart yang ditemukan
	if err := tx.Delete(&cart).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit transaksi
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response setelah cart berhasil dihapus
	c.JSON(http.StatusOK, model.CartResponse{
		Message: "Cart deleted successfully",
		Data:    cartDeleteResponse,
	})
}
