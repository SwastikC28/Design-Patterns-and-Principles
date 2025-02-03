package model

import "errors"

type VendingMachineBank interface {
	AddBalance(amount int) error
	Refund(amount int)
	GetBalance() int
}

func NewVendingMachineBank() VendingMachineBank {
	return &vendingMachineBank{
		balance: 0,
	}
}

type vendingMachineBank struct {
	balance int
}

func (bank *vendingMachineBank) AddBalance(amount int) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}

	bank.balance += amount
	return nil
}

func (bank *vendingMachineBank) Refund(amount int) {
	bank.balance -= amount
}

func (bank *vendingMachineBank) GetBalance() int {
	return bank.balance
}
