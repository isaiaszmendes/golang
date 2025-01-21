package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type NotificationService interface {
	SendNotification(message string)
}

// interface da abstract factory
type BillingFactory interface {
	CreatePaymentProcessor() PaymentProcessor
	CreateNotificationService() NotificationService
}

// implementaçoes concretas de pagamento do mercado pago

type MercadoPagoProcessor struct{}

func (m *MercadoPagoProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Pagando com MercadoPago: $ %.2f \n", amount)
}

type MercadoPagoNotification struct{}

func (m *MercadoPagoNotification) SendNotification(message string) {
	fmt.Printf("Enviando notificação com MercadoPago: %s\n", message)
}

// implementaçoes concretas de pagamento do pagarme

type PagarMeProcessor struct{}

func (p *PagarMeProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Pagando com PagarMe: $ %.2f \n", amount)
}

type PagarMeNotification struct{}

func (p *PagarMeNotification) SendNotification(message string) {
	fmt.Printf("Enviando notificação com PagarMe: %s\n", message)
}

// implementaçoes concretas da abstract factory

type MercadoPagoFactory struct{}

func (m *MercadoPagoFactory) CreatePaymentProcessor() PaymentProcessor {
	return &MercadoPagoProcessor{}
}

func (m *MercadoPagoFactory) CreateNotificationService() NotificationService {
	return &MercadoPagoNotification{}
}

type PagarMeFactory struct{}

func (p *PagarMeFactory) CreatePaymentProcessor() PaymentProcessor {
	return &PagarMeProcessor{}
}

func (p *PagarMeFactory) CreateNotificationService() NotificationService {
	return &PagarMeNotification{}
}

func ProcessPayment(factory BillingFactory, amount float64, message string) {
	paymentProcessor := factory.CreatePaymentProcessor()
	paymentProcessor.ProcessPayment(amount)

	notificationService := factory.CreateNotificationService()
	notificationService.SendNotification(message)
}

func main() {
	fmt.Println("Usando o pagarme")
	pagarmeFactory := &PagarMeFactory{}
	ProcessPayment(pagarmeFactory, 100, "Pagamento de 100 reais")

	fmt.Println("Usando o mercado pago")
	mercadoPagoFactory := &MercadoPagoFactory{}
	ProcessPayment(mercadoPagoFactory, 200, "Pagamento de 200 reais")
}

	