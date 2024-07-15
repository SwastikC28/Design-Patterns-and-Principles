package creditcard

import (
	"fmt"
	"paymentfactory/utils"
)

func NewCreditCard() utils.Payment {
	return &CreditCard{}
}

type CreditCard struct{}

func (creditCard *CreditCard) Initialize() {
	fmt.Println("Initializing Payment through credit card")
}
func (creditCard *CreditCard) Payment() {
	fmt.Println("Action Payment of credit card")
}
