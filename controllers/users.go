package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sirapo/models"
)

// READ
func FindUsers(c *gin.Context) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch users",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

// CREATE
type CreateUsersInput struct {
	Username string      `json:"username" binding:"required"`
	Email    string      `json:"email" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Role     models.Role `json:"role" binding:"required"`
}

func CreateUsers(c *gin.Context) {
	// Validate input
	var input CreateUsersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	// Create user
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}
	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	// Save to database
	result := models.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

// DELETE
func DeleteUsers(c *gin.Context) {
	var users models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&users)
	c.JSON(http.StatusOK, gin.H{"data": true, "message": "User deleted successfully"})
}

type UpdateBookInput struct {
	Username string      `json:"username" binding:"required"`
	Email    string      `json:"email" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Role     models.Role `json:"role" binding:"required"`
}

// UPDATE
func UpdateUsers(c *gin.Context) {
	// Get model if exist
	var users models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users.Username = input.Username
	users.Email = input.Email
	users.Password = input.Password
	users.Role = input.Role

	// Save updates to database
	if err := models.DB.Save(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users, "message": "User updated successfully"})
}
