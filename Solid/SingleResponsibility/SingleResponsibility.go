package main

import "fmt"

type Account struct {
	accNumber uint
	name      string
	balance   uint
}

func (acc *Account) Deposit(amount uint) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount uint) {
	acc.balance -= amount
}

func (acc *Account) BalanceEnquiry() uint {
	return acc.balance
}

func NewAccount(name string) *Account {
	return &Account{
		accNumber: 1,
		name:      name,
		balance:   0,
	}
}

func main() {
	acc := NewAccount("Swastik")
	acc.Deposit(10)
	acc.Withdraw(5)

	fmt.Println(acc.BalanceEnquiry())
}
