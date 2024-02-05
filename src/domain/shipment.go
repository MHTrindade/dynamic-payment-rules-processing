package domain

import "fmt"

type Shipment struct {
	ProductDescription string
}

func NewShipment(productDescription string) *Shipment {
	return &Shipment{ProductDescription: productDescription}
}

func (s *Shipment) Generate() error {
	fmt.Printf("Shipping label generated for the product: %s\n", s.ProductDescription)
	return nil
}
