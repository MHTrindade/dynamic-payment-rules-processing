package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"payment.com/payment-rules/src/domain"
)

func ProcessPayment(c *gin.Context) {
	var request domain.PaymentRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	processor := createPaymentProcessor(request.Product.ProductType)

	transactionID, err := processor.ProcessPayment(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction_id": transactionID})
}

func createPaymentProcessor(productType domain.ProductType) *domain.PaymentProcessor {
	var strategy domain.PaymentStrategy

	switch productType {
	case domain.CommonProductType:
		strategy = &domain.CommonPaymentStrategy{}
	case domain.PhysicalProductType:
		strategy = &domain.PhysicalProductPaymentStrategy{}
	case domain.BookProductType:
		strategy = &domain.BookPaymentStrategy{}
	case domain.VideoProductType:
		strategy = &domain.VideoPaymentStrategy{}
	default:
		strategy = &domain.CommonPaymentStrategy{}
	}

	return &domain.PaymentProcessor{Strategy: strategy}
}
