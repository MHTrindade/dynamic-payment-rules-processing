package domain

import "fmt"

type VideoPaymentStrategy struct{}

func (s *VideoPaymentStrategy) ProcessPayment(request PaymentRequest) (string, error) {
	fmt.Println("Processando pagamento para vídeo...")

	shipment := NewShipment(request.Product.Description)
	err := shipment.Generate()
	if err != nil {
		return "", fmt.Errorf("erro ao gerar guia de remessa: %s", err.Error())
	}

	if request.Product.RequiresFirstAidVideo {
		product := Product{Description: "video gratuito"}
		shipment := NewShipment(product.Description)
		err := shipment.Generate()
		if err != nil {
			return "", fmt.Errorf("erro ao gerar guia de remessa: %s", err.Error())
		}
	}

	return "video_transaction_id_123", nil
}
