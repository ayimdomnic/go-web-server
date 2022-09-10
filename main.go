package main

import (
	"net/http"

	"github.com/ayimdomnic/go-web-server/controllers"
	"github.com/ayimdomnic/go-web-server/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to MegaPesa Apis",
		})
	})
	r.GET("/lottos", controllers.FindLottos)
	r.POST("/register", controllers.Register)

	r.Run(":3000")
}
