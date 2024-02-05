package infra

import "fmt"

// Email é uma estrutura para simular o envio de e-mails
type Email struct{}

// NewEmail cria uma nova instância de Email
func NewEmail() *Email {
	return &Email{}
}

// Send simula o envio de e-mails
func (e *Email) Send(subject string) error {
	fmt.Printf("Assunto: %s\n", subject)
	// Lógica real de envio de e-mails aqui
	return nil
}
