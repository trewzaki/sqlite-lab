package controllers

import (
	"sqlite-lab/models"

	"github.com/gin-gonic/gin"
)

// CreateUser controller function TODO
func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
	return
}

// UpdateUser controller function TODO
func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
	return
}

// GetUserByID controller function TODO
func GetUserByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
	return

}

// GetAllUser controller function
func GetAllUser(c *gin.Context) {
	userModel := models.User{}

	userModels, getAllErr := userModel.GetAll()
	if getAllErr != nil {
		c.JSON(200, gin.H{
			"message": getAllErr,
			"success": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"model":   userModels,
		"success": true,
	})
	return
}

// DeleteUser controller function TODO
func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
	return
}

// UnDeleteUser controller function TODO
func UnDeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
	return
}
