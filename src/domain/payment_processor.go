package domain

type PaymentProcessor struct {
	Strategy PaymentStrategy
}

func (p *PaymentProcessor) ProcessPayment(request PaymentRequest) (string, error) {
	return p.ProcessPayment(request)
}
