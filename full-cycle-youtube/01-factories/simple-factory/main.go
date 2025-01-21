package main

import "fmt"

type PaymentProcessor interface {
	Pay()
}

type PagarMe struct{}

func (p *PagarMe) Pay() {
	fmt.Println("Pagando com PagarMe")
}

type MercadoPago struct{}

func (m *MercadoPago) Pay() {
	fmt.Println("Pagando com MercadoPago")
}

func SimplePaymentFactory(provider string) PaymentProcessor {
	switch provider {
	case "pagarme":
		return &PagarMe{}
	case "mercadopago":
		return &MercadoPago{}
	default:
		return nil
	}
}

func main() {
	paymentProcessor := SimplePaymentFactory("mercadopago")
	paymentProcessor.Pay()
}
