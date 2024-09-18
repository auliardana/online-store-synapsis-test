package controller

import (
	"fmt"
	"log"
	"net/http"
	"online-store/app/db"
	"online-store/app/model"
	"os"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"

	// "github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	// "github.com/midtrans/midtrans-go/iris"

	"github.com/gin-gonic/gin"
)

// Create is a method to create a role by json request body
// @Summary Create a payment order
// @Description Create a role with its name and its optional description, worksets, resources, and services. Worksets, resources, and services can be added to user with UUID/UUIDs
// @Tags Order
// @Accept  json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer 'Add access token here')
// @Param   orders body model.OrderRequest true "Create Order"
// @Success 201 {object} model.OrderResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /order [post]
func CheckoutOrder(c *gin.Context) {
	var orderRequest model.OrderRequest

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
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

	// Parse product UUID
	productUUID, err := uuid.Parse(orderRequest.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product UUID"})
		return
	}

	// Get product price
	var product model.Product
	if err := db.DBConn.Where("id = ?", productUUID).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	// Calculate total price
	totalPrice := orderRequest.Quantity * product.Price

	order := model.Order{
		ProductID:     productUUID,
		Quantity:      orderRequest.Quantity,
		UserID:        userUUID,
		TotalPrice:    totalPrice,
		PaymentStatus: "unpaid",
	}

	tx := db.DBConn.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create payment request to Midtrans
	snapURL, err := CreateMidtransTransaction(&order)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment request"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Order created successfully",
		"data":        order,
		"payment_url": snapURL,
	})
}

// CreateMidtransTransaction creates a new transaction with Midtrans and returns the payment URL
func CreateMidtransTransaction(order *model.Order) (string, error) {
	// Use order.ID directly since it's already of type uuid.UUID
	orderUUID := order.ID

	// Initiate Snap client
	s := snap.Client{}
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		return "", fmt.Errorf("MIDTRANS_SERVER_KEY is not set")
	}
	s.New(serverKey, midtrans.Sandbox) // Change to midtrans.Production for real transactions

	// Find user data from order
	var user model.User
	if err := db.DBConn.Where("id = ?", order.UserID).First(&user).Error; err != nil {
		return "", err
	}

	// Calculate total price
	totalPrice := order.Quantity * order.TotalPrice

	// Create Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderUUID.String(),
			GrossAmt: int64(totalPrice), // Midtrans expects integer value for amount
		},
		Gopay: &snap.GopayDetails{
			EnableCallback: true,
			CallbackUrl:    "https://a518-101-128-100-16.ngrok-free.app/midtrans/payment-callback",
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.FirstName,
			LName: user.LastName,
			Email: user.Email,
			Phone: user.Phone,
		},
	}

	// Execute request to create Snap transaction to Midtrans Snap API
	snapResp, err := s.CreateTransaction(req)
	if err != nil {
		return "", err
	}
	return snapResp.RedirectURL, nil
}

type PaymentCallbackRequest struct {
	OrderID           string `json:"order_id"`
	TransactionID     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	SignatureKey      string `json:"signature_key"`
}

func PaymentCallbackHandler(c *gin.Context) {
	var req PaymentCallbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil order berdasarkan OrderID
	var order model.Order
	if err := db.DBConn.Where("id = ?", req.OrderID).First(&order).Error; err != nil {
		log.Println("Order not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Validasi status pembayaran
	if req.TransactionStatus == "success" {
		// Mulai transaksi database
		tx := db.DBConn.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// Update status order menjadi "paid"
		order.PaymentStatus = "paid"
		if err := tx.Save(&order).Error; err != nil {
			tx.Rollback()
			log.Println("Failed to update order status:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
			return
		}

		// Ambil produk terkait
		var product model.Product
		if err := tx.Where("id = ?", order.ProductID).First(&product).Error; err != nil {
			tx.Rollback()
			log.Println("Product not found:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// Kurangi stok produk
		product.Stock -= order.Quantity
		if product.Stock < 0 {
			tx.Rollback()
			log.Println("Negative stock detected")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Negative stock"})
			return
		}

		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			log.Println("Failed to update product stock:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
			return
		}

		// Commit transaksi
		if err := tx.Commit().Error; err != nil {
			log.Println("Transaction commit failed:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
			return
		}

		// Respon sukses
		c.JSON(http.StatusOK, gin.H{"message": "Payment processed and stock updated"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Payment not successful"})
	}
}
