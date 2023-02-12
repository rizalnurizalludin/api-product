package controllers

import (
	"net/http"

	"github.com/rizalnurizalludin/api-product/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type UpdateProductInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func Index(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []entity.Product
	db.Find(&products)

	if len(products) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Record Not Found", "message": "Product Tidak Ada"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": products})
	}
}

func Create(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create todo item
	product := entity.Product{Name: input.Name, Price: input.Price}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Show(c *gin.Context) {
	var product entity.Product

	db := c.MustGet("db").(*gorm.DB)
	data := db.Where("id = ?", c.Param("id")).First(&product).Error
	if data != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Record Not Found", "message": "Product with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Update(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get entity if exist
	var product entity.Product
	data := db.Where("id = ?", c.Param("id")).First(&product).Error
	if data != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Record Not Found", "message": "Product with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	// Validate input
	var input UpdateProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput entity.Product
	updatedInput.Name = input.Name
	updatedInput.Price = input.Price

	db.Model(&product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Delete(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product entity.Product
	data := db.Where("id = ?", c.Param("id")).First(&product).Error
	if data != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Record Not Found", "message": "Product with ID " + c.Param("id") + " Not Found", "data": false})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "message": "Deleted Success", "data": true})
}
