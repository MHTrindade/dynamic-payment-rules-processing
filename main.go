package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"payment.com/payment-rules/src/application"
	"payment.com/payment-rules/src/domain"
)

// Simula a ativação de uma nova associação de membro
func activateNewMembership(description string) error {
	fmt.Printf("Associação ativada para novo membro: %s\n", description)
	// Lógica real de ativação de nova associação aqui
	return nil
}

// Simula a aplicação de upgrade em uma associação de membro
func applyMembershipUpgrade(description string) error {
	fmt.Printf("Upgrade aplicado na associação de membro: %s\n", description)
	// Lógica real de aplicação de upgrade aqui
	return nil
}

// Simula o envio de e-mail de boas-vindas
func sendWelcomeEmail(description string) error {
	fmt.Printf("E-mail de boas-vindas enviado para: %s\n", description)
	// Lógica real de envio de e-mail de boas-vindas aqui
	return nil
}

// Simula o envio de e-mail de confirmação de upgrade
func sendUpgradeConfirmationEmail(description string) error {
	fmt.Printf("E-mail de confirmação de upgrade enviado para: %s\n", description)
	// Lógica real de envio de e-mail de confirmação de upgrade aqui
	return nil
}

// Simula a adição de um vídeo gratuito de primeiros socorros à guia de remessa
func addFreeFirstAidVideoToShippingLabel() error {
	fmt.Println("Vídeo gratuito de primeiros socorros adicionado à guia de remessa")
	// Lógica real de adição de vídeo gratuito de primeiros socorros à guia de remessa aqui
	return nil
}

// Manipulador para endpoint de pagamento usando Gin
func handlePayment(c *gin.Context) {
	// Decodifica o corpo JSON da solicitação
	var paymentRequest domain.PaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON da solicitação"})
		return
	}

	// Processa o pagamento
	paymentProcessor := &domain.PaymentProcessor{}
	transactionID, err := paymentProcessor.ProcessPayment(paymentRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Erro ao processar o pagamento: %s", err.Error())})
		return
	}

	// Responde com o ID da transação
	c.JSON(http.StatusOK, gin.H{"transaction_id": transactionID})
}

func main() {
	r := gin.Default()

	application.SetupRoutes(r)

	r.Run(":8080")
}
