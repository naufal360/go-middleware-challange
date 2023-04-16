package controllers

import (
	"go-middleware-challange/database"
	"go-middleware-challange/helpers"
	"go-middleware-challange/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userId := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userId

	if err := db.Debug().Create(&Product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Product)
}

func GetAllProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := []models.Product{}
	userId := uint(userData["id"].(float64))
	isAdmin := userData["admin"].(bool)

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	if !isAdmin {
		if err := db.Model(&Product).Where("user_id = ?", userId).Find(&Product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := db.Model(&Product).Find(&Product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, Product)
}

func GetProductById(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	// userId := uint(userData["id"].(float64))
	// isAdmin := userData["admin"].(bool)

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	if err := db.Model(&Product).Where("id = ?", productId).First(&Product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userId := uint(userData["id"].(float64))
	isAdmin := userData["admin"].(bool)

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	if !isAdmin {
		Product.UserID = userId
	}
	Product.ID = uint(productId)

	if err := db.Model(&Product).Where("id = ?", productId).Updates(
		models.Product{
			Title:       Product.Title,
			Description: Product.Description,
		},
	).Scan(&Product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func DeleteProductById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userId := uint(userData["id"].(float64))
	isAdmin := userData["admin"].(bool)

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	if !isAdmin {
		Product.UserID = userId
	}
	Product.ID = uint(productId)

	if err := db.Model(&Product).Where("id = ?", productId).Delete(&Product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Product",
	})
}
