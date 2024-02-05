package infra

import "fmt"

type Email struct{}

func NewEmail() *Email {
	return &Email{}
}

func (e *Email) Send(subject string) error {
	fmt.Printf("Assunto: %s\n", subject)
	return nil
}
