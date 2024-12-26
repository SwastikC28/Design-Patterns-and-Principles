package adapter

import "fmt"

// PaymentProcessor interface
type PaymentProcessor interface {
	ProcessPayment(int, string) (bool, error)
}

// Stripe Adapter
type stripeClient struct{}

func (client *stripeClient) makeStripePayment(totalAmount int, paymentMethod string, customerDetails map[string]interface{}) (bool, error) {
	return true, nil
}

type stripePaymentProcessorAdapter struct {
	client *stripeClient
}

func NewStripePaymentProcessorAdapter() PaymentProcessor {
	return &stripePaymentProcessorAdapter{
		client: &stripeClient{},
	}
}

func (adapter *stripePaymentProcessorAdapter) ProcessPayment(amount int, currency string) (bool, error) {
	success, err := adapter.client.makeStripePayment(amount, "card", map[string]interface{}{"customerID": "12345"})
	if err != nil {
		return false, fmt.Errorf("Stripe Payment failed: %v", err)
	}
	return success, nil
}

// PayPal Adapter
type paypalClient struct{}

func (client *paypalClient) processPaypalPayment(amount int, currency string) (bool, error) {
	// Example: Add actual logic for PayPal payment here
	return true, nil
}

type paypalPaymentProcessorAdapter struct {
	client *paypalClient
}

func NewPaypalPaymentProcessorAdapter() PaymentProcessor {
	return &paypalPaymentProcessorAdapter{
		client: &paypalClient{},
	}
}

func (adapter *paypalPaymentProcessorAdapter) ProcessPayment(amount int, currency string) (bool, error) {
	success, err := adapter.client.processPaypalPayment(amount, currency)
	if err != nil {
		return false, fmt.Errorf("PayPal Payment failed: %v", err)
	}
	return success, nil
}

// Razorpay Adapter
type razorpayClient struct{}

func (client *razorpayClient) initiateRazorpayTransaction(amount int, customerDetails map[string]interface{}) (bool, error) {
	// Example: Add actual logic for Razorpay payment here
	return true, nil
}

type razorpayPaymentProcessorAdapter struct {
	client *razorpayClient
}

func NewRazorpayPaymentProcessorAdapter() PaymentProcessor {
	return &razorpayPaymentProcessorAdapter{
		client: &razorpayClient{},
	}
}

func (adapter *razorpayPaymentProcessorAdapter) ProcessPayment(amount int, currency string) (bool, error) {
	success, err := adapter.client.initiateRazorpayTransaction(amount, map[string]interface{}{"customerID": "67890"})
	if err != nil {
		return false, fmt.Errorf("Razorpay Payment failed: %v", err)
	}
	return success, nil
}
