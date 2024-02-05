package domain

import "fmt"

type RoyaltyShipment struct {
	Description string
}

func NewRoyaltyShipment(description string) *RoyaltyShipment {
	return &RoyaltyShipment{
		Description: description,
	}
}

func (rs *RoyaltyShipment) Create() error {
	fmt.Printf("Remessa criada para o departamento de royalties para o livro: %s\n", rs.Description)
	return nil
}
