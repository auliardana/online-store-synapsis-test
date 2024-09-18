package controller

import (
	"net/http"
	"online-store/app/db"
	"online-store/app/model"

	"github.com/gin-gonic/gin"
)

// Create is a method to create category
// @Summary create category
// @Description create category
// @Tags Category
// @Accept  json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer 'Add access token here')
// @Param   productCategory body model.ProductCategoryRequest true "cart need name of category"
// @Success 201 {object} model.ProductCategoryResponse	"created successfully"
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /category [post]
func CreateCategory(c *gin.Context) {
	var category model.ProductCategory
	var categoryRequest model.ProductCategoryRequest
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category = model.ProductCategory{
		Name: categoryRequest.Name,
	}

	tx := db.DBConn.Begin()

	if err := tx.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
		"data":    category,
	})

}

// GetAllCategories is a method to get all categories
// @Summary view product list by product category
// @Description Get a list of all product categories
// @Tags Category
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer 'Add access token here')
// @Success 200 {object} model.ProductCategoryResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /category [get]
func GetAllCategories(c *gin.Context) {
	var categories []model.ProductCategory

	if err := db.DBConn.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find categories.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
