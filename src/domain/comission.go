package domain

import "fmt"

type Commission struct{}

func NewCommission() *Commission {
	return &Commission{}
}

func (c *Commission) ProcessPayment(amount float64) error {
	fmt.Printf("Pagamento de comissão processado: %f\n", amount)
	return nil
}
