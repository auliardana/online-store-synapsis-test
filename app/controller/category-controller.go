package controller

import (
	"net/http"
	"online-store/app/db"
	"online-store/app/model"
	"github.com/gin-gonic/gin"

)

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