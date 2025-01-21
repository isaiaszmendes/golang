package main

import "fmt"

// interface do processador
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}



// implementaçoes concretas

type MercadoPago struct{}

func (m *MercadoPago) ProcessPayment(amount float64) {
	println("Pagando com MercadoPago: ", amount)
}

type PagarMe struct{}

func (p *PagarMe) ProcessPayment(amount float64) {
	println("Pagando com PagarMe: ", amount)
}

// interface da fabrica
type PaymentFactory interface {
	CreatePaymentProcessor() PaymentProcessor
}

// implementaçoes concretas da fabrica
type MercadoPagoFactory struct{}

func (m *MercadoPagoFactory) CreatePaymentProcessor() PaymentProcessor {
	return &MercadoPago{}
}

type PagarMeFactory struct{}

func (p *PagarMeFactory) CreatePaymentProcessor() PaymentProcessor {
	return &PagarMe{}
}

// Esse é o método responsável por processar o pagamento
func ProcessPayment(factory PaymentFactory, amount float64) {
	paymentProcessor := factory.CreatePaymentProcessor()
	paymentProcessor.ProcessPayment(amount)
}

func main() {
	fmt.Println("Usando o pagarme")
	pagarmeFactory := &PagarMeFactory{}
	ProcessPayment(pagarmeFactory, 100)

	fmt.Println("Usando o mercado pago")
	mercadoPagoFactory := &MercadoPagoFactory{}
	ProcessPayment(mercadoPagoFactory, 200)

}