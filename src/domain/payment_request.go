package domain

type PaymentRequest struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Product    Product `json:"product"`
	MemberType string  `json:"member_type"`
}
