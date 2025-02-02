package model

import (
	"errors"
)

type PaymentMode interface {
	AddBalance(amount int)
	GetBalance() int
	DeductBalance(amount int) error
}

type CreditCard struct {
	balance int
}

func NewCreditCard() PaymentMode {
	return &CreditCard{
		balance: 0,
	}
}

func (c *CreditCard) AddBalance(amount int) {
	c.balance += amount
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

func (c *CreditCard) DeductBalance(amount int) error {
	if c.balance < amount {
		return errors.New("insufficient balance")
	}

	c.balance -= amount
	return nil
}

func (c *CreditCard) CreditBalance(amount int) {
	c.balance += amount
}

type Cash struct {
	balance int
}

func NewCashWallet() PaymentMode {
	return &Cash{
		balance: 0,
	}
}

func (c *Cash) AddBalance(amount int) {
	c.balance += amount
}

func (c *Cash) GetBalance() int {
	return c.balance
}

func (c *Cash) DeductBalance(amount int) error {
	if c.balance < amount {
		return errors.New("insufficient balance")
	}

	c.balance -= amount
	return nil
}

func (c *Cash) CreditBalance(amount int) {
	c.balance += amount
}
