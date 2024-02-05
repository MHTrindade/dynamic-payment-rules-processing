package domain

import (
	"fmt"
)

type PhysicalProductPaymentStrategy struct{}

func (s *PhysicalProductPaymentStrategy) ProcessPayment(request PaymentRequest) (string, error) {
	fmt.Println("Proceasando pagamento para produto físico...")

	shipment := NewShipment(request.Product.Description)
	err := shipment.Generate()
	if err != nil {
		return "", fmt.Errorf("erro ao gerar guia de remessa: %s", err.Error())
	}

	commission := NewCommission()
	err = commission.ProcessPayment(request.Amount)
	if err != nil {
		return "", fmt.Errorf("erro ao processar o pagamento da comissão: %s", err.Error())
	}

	return "physical_product_transaction_id_123", nil
}
