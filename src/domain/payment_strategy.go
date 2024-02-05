package domain

type PaymentStrategy interface {
	ProcessPayment(request PaymentRequest) (string, error)
}
