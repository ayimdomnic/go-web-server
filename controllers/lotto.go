package controllers

import (
	"net/http"

	"github.com/ayimdomnic/go-web-server/models"
	"github.com/gin-gonic/gin"
)

// GET all Lotto Entries
// GET /lotto

func FindLottos(c *gin.Context) {
	var lottos []models.Lotto
	models.DB.Find(&lottos)

	c.JSON(http.StatusOK, gin.H{"data": lottos})
}
