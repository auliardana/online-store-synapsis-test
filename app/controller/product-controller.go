package controller

import (
	"fmt"
	"net/http"
	"online-store/app/model"
	"strconv"

	"online-store/app/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllProducts(c *gin.Context) {
	// var initialUsers []model.Product

	type paginate struct {
		Page     int
		Limit    int
		Category uuid.UUID
	}

	// Default pagination values
	defaultQueryValues := paginate{
		Page:     1,
		Limit:    10,
		Category: uuid.Nil,
	}

	// Get query params from URL
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "page":
			page, err := strconv.Atoi(queryValue)
			if err == nil && page >= 1 {
				defaultQueryValues.Page = page
			}
		case "limit":
			limit, err := strconv.Atoi(queryValue)
			if err == nil {
				if limit >= 10 && limit <= 100 {
					defaultQueryValues.Limit = limit
				}
			}
		case "category":
			categoryUUID, err := uuid.Parse(queryValue)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid category UUID",
				})
				return
			}
			defaultQueryValues.Category = categoryUUID

			// Validate category
			var product model.Product
			fmt.Println(defaultQueryValues.Category)
			if err := db.DBConn.Where("category_id = ?", defaultQueryValues.Category).First(&product).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Category not found ",
				})
				return
			}
		}
	}

	// Count total products for pagination
	var totalProducts int64
	countQuery := db.DBConn.Model(&model.Product{})
	if defaultQueryValues.Category != uuid.Nil {
		countQuery = countQuery.Where("category_id = ?", defaultQueryValues.Category)
	}
	if err := countQuery.Count(&totalProducts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't count products.",
			"error":   err.Error(),
		})
		return
	}

	if totalProducts == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No products found.",
		})
		return
	}

	// Calculate total pages
	totalPages := int((totalProducts + int64(defaultQueryValues.Limit) - 1) / int64(defaultQueryValues.Limit))
	offset := (defaultQueryValues.Page - 1) * defaultQueryValues.Limit

	// Get products with pagination
	var products []model.Product
	queryProduct := db.DBConn.Limit(defaultQueryValues.Limit).Offset(offset)

	if defaultQueryValues.Category != uuid.Nil {
		queryProduct = queryProduct.Where("category_id = ?", defaultQueryValues.Category)
	}

	if err := queryProduct.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find products.",
			"error":   err.Error(),
		})
		return
	}

	// Return result
	c.JSON(http.StatusOK, gin.H{
		"products":    products,
		"totalPage":   totalPages,
		"currentPage": defaultQueryValues.Page,
		"totalItems":  totalProducts,
	})
}

func CreateProduct(c *gin.Context) {
	var productRequest model.ProductRequest
	var product model.Product

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert CategoryID from string to uuid.UUID
	categoryUUID, err := uuid.Parse(productRequest.Category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category UUID"})
		return
	}

	product = model.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
		ImageURL:    productRequest.ImageURL,
		CategoryID:  categoryUUID,
	}

	tx := db.DBConn.Begin()

	if err := tx.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, model.ProductResponse{
		Message: "Create Product successfull",
		Data:    productRequest,
	})
}
