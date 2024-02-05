package domain

import "fmt"

type BookPaymentStrategy struct{}

func (s *BookPaymentStrategy) ProcessPayment(request PaymentRequest) (string, error) {
	fmt.Println("Processando pagamento para livro...")

	commission := NewCommission()
	err := commission.ProcessPayment(request.Amount)
	if err != nil {
		return "", fmt.Errorf("erro ao processar o pagamento da comissão: %s", err.Error())
	}

	royaltyShipment := NewRoyaltyShipment(request.Product.Description)
	err = royaltyShipment.Create()
	if err != nil {
		return "", fmt.Errorf("erro ao criar remessa para o departamento de royalties: %s", err.Error())
	}

	return "book_transaction_id_123", nil
}
