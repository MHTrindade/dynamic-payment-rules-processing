package application

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/processPayment", ProcessPayment)
}
