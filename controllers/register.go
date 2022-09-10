package controllers

import (
	"net/http"

	"github.com/ayimdomnic/go-web-server/models"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" bindiing:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	u.FirstName = input.FirstName
	u.LastName = input.LastName
	u.Email = input.Email
	u.PhoneNumber = input.PhoneNumber

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered User Successfully"})
}
