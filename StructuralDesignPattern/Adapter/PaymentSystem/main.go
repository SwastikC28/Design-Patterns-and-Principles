package main

import (
	"fmt"
	"paymentprocessor/adapter"
)

func main() {
	fmt.Println("Payment Gateway Adapter")

	// PayPal Payment
	paypalProcessor := adapter.NewPaypalPaymentProcessorAdapter()
	_, err := paypalProcessor.ProcessPayment(2000, "INR")
	if err != nil {
		fmt.Println("Error processing payment through PayPal:", err)
	} else {
		fmt.Println("Payment Successful through PayPal Payment Processor")
	}

	// Stripe Payment
	stripeProcessor := adapter.NewStripePaymentProcessorAdapter()
	_, err = stripeProcessor.ProcessPayment(2000, "INR")
	if err != nil {
		fmt.Println("Error processing payment through Stripe:", err)
	} else {
		fmt.Println("Payment Successful through Stripe Payment Processor")
	}

	// Razorpay Payment
	razorpayProcessor := adapter.NewRazorpayPaymentProcessorAdapter()
	_, err = razorpayProcessor.ProcessPayment(2000, "INR")
	if err != nil {
		fmt.Println("Error processing payment through Razorpay:", err)
	} else {
		fmt.Println("Payment Successful through Razorpay Payment Processor")
	}
}
