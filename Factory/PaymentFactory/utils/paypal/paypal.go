package paypal

import (
	"fmt"
	"paymentfactory/utils"
)

func NewPaypal() utils.Payment {
	return &Paypal{}
}

type Paypal struct{}

func (paypal *Paypal) Initialize() {
	fmt.Println("Initializing Payment through paypal")
}
func (paypal *Paypal) Payment() {
	fmt.Println("Action Payment paypal")
}
