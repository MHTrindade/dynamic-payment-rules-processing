package domain

// PaymentStrategy é uma interface para estratégias de pagamento
type PaymentStrategy interface {
	ProcessPayment(request PaymentRequest) (string, error)
}
