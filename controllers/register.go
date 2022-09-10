package controllers

import (
	"net/http"

	"github.com/ayimdomnic/go-web-server/models"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
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

	c.JSON(http.StatusOK, gin.H{"message": "Register User"})
}
