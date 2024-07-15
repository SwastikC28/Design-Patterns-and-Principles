package factory

import (
	"errors"
	"paymentfactory/utils"
	"paymentfactory/utils/bitcoin"
	"paymentfactory/utils/creditcard"
	"paymentfactory/utils/paypal"
)

var registry map[string]utils.Payment

func init() {
	registry = make(map[string]utils.Payment)
	registry["cc"] = creditcard.NewCreditCard()
	registry["paypal"] = paypal.NewPaypal()
	registry["bitcoin"] = bitcoin.NewBitcoin()
}

type PaymentFactory struct {
}

func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{}
}

func (factory *PaymentFactory) NewPayment(paymentType string) (utils.Payment, error) {
	paymentMethod, ok := registry[paymentType]
	if !ok {
		return nil, errors.New("unsupported payment type")
	}

	return paymentMethod, nil
}
