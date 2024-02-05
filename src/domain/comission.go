// No arquivo domain/commission_processor.go

package domain

import "fmt"

// Commission é uma estrutura para processar pagamentos de comissão
type Commission struct{}

// NewCommission cria uma nova instância de Commission
func NewCommission() *Commission {
	return &Commission{}
}

// ProcessPayment processa o pagamento de comissão
func (c *Commission) ProcessPayment(amount float64) error {
	fmt.Printf("Pagamento de comissão processado: %f\n", amount*0.1) // 10% de comissão
	// Lógica real de pagamento de comissão aqui
	return nil
}
