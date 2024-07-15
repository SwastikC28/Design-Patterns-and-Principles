package bitcoin

import (
	"fmt"
	"paymentfactory/utils"
)

func NewBitcoin() utils.Payment {
	return &Bitcoin{}
}

type Bitcoin struct{}

func (bitcoin *Bitcoin) Initialize() {
	fmt.Println("Initializing Payment through bitcoin")
}
func (bitcoin *Bitcoin) Payment() {
	fmt.Println("Action Payment bitcoin")
}
