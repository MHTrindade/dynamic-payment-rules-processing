package domain

// PaymentRequest é uma estrutura que representa uma solicitação de pagamento
type PaymentRequest struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Product    Product `json:"product"`
	MemberType string  `json:"member_type"` // Tipo de associação (novo, upgrade, etc.)
}
