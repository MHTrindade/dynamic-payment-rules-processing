package domain

import (
	"fmt"
)

type CommonPaymentStrategy struct{}

func (s *CommonPaymentStrategy) ProcessPayment(request PaymentRequest) (string, error) {
	fmt.Println("Processamento comum de pagamento...")
	return "common_transaction_id_123", nil
}
