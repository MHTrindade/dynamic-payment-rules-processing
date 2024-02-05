package main

import (
	"github.com/gin-gonic/gin"
	"payment.com/payment-rules/src/application"
)

func main() {
	r := gin.Default()

	application.SetupRoutes(r)

	r.Run(":8080")
}
