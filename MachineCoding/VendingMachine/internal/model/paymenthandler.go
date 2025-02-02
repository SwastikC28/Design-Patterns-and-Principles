package model

import "fmt"

type PaymentHandler interface {
	Pay(paymentMode PaymentMode, amount int) error
	Refund(paymentMode PaymentMode, amount int) error
}

func NewPaymentHandler(vendingMachineBank VendingMachineBank) PaymentHandler {
	return &paymentHandler{
		creditCardPaymentHandler: &CreditCardPaymentHandler{vendingMachineBank: vendingMachineBank},
	}
}

type paymentHandler struct {
	creditCardPaymentHandler *CreditCardPaymentHandler
}

func (p *paymentHandler) Pay(paymentMode PaymentMode, amount int) error {
	switch paymentMode.(type) {
	case *CreditCard:
		return p.creditCardPaymentHandler.Pay(paymentMode, amount)
	default:
		return fmt.Errorf("invalid payment mode")
	}
}

func (p *paymentHandler) Refund(paymentMode PaymentMode, amount int) error {
	switch paymentMode.(type) {
	case *CreditCard:
		p.creditCardPaymentHandler.Refund(paymentMode, amount)
		return nil
	default:
		return fmt.Errorf("invalid payment mode")
	}
}

// Use factory here
type CreditCardPaymentHandler struct {
	vendingMachineBank VendingMachineBank
}

func (c *CreditCardPaymentHandler) Pay(paymentMode PaymentMode, amount int) error {
	fmt.Println("Connecting to credit cards servers")

	err := paymentMode.DeductBalance(amount)
	if err != nil {
		return err
	}

	c.vendingMachineBank.AddBalance(amount)
	fmt.Println("Payment successful")
	return nil
}

func (c *CreditCardPaymentHandler) Refund(paymentMode PaymentMode, amount int) {
	fmt.Println("Connecting to credit cards servers")

	c.vendingMachineBank.Refund(amount)
	paymentMode.AddBalance(amount)

	fmt.Println("Refund successful")

}
