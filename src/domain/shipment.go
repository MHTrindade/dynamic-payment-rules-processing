// Adicione esta struct no arquivo domain/shipment.go

package domain

import "fmt"

// Shipment is a structure that represents the action of generating a shipping label
type Shipment struct {
	ProductDescription string
	// Add other fields relevant to the shipping label
}

// NewShipment creates a new instance of Shipment
func NewShipment(productDescription string) *Shipment {
	return &Shipment{ProductDescription: productDescription}
}

// Generate generates the shipping label
func (s *Shipment) Generate() error {
	fmt.Printf("Shipping label generated for the product: %s\n", s.ProductDescription)
	// Real logic for generating a shipping label goes here
	return nil
}
