package main

import (
	"fmt"
	"paymentfactory/utils/factory"
)

func main() {
	paymentFactory := factory.NewPaymentFactory()
	paymentGateway, err := paymentFactory.NewPayment("cc")
	if err != nil {
		fmt.Println("Error creating payment:", err)
		return
	}

	paymentGateway.Initialize()
	paymentGateway.Payment()
}
