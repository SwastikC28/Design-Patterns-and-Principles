package main

import (
	"fmt"
	"vendingmachine/internal/model"
)

func main() {
	fmt.Println("Vending Machine LLD")

	user := model.NewUser(false)
	creditCard := model.NewCreditCard()
	cashWallet := model.NewCashWallet()

	user.AddPaymentMode(model.CREDIT_CARD, creditCard)
	user.AddPaymentMode(model.CASH, cashWallet)

	lays := model.NewChip("Lays", 10)
	oreo := model.NewBiscuit("oreo", 20)
	cola := model.NewDrink("Thumbs Up", 40)

	admin := model.NewUser(true)
	inventory := model.NewInventory()
	inventory.Restock(admin, lays, 10)
	inventory.Restock(admin, oreo, 10)
	inventory.Restock(admin, cola, 10)

	vendingMachineBank := model.NewVendingMachineBank()
	vendingMachine := model.NewVendingMachine(inventory, vendingMachineBank)

	vendingMachine.DisplayProducts()
	err := vendingMachine.BuyProduct("Lays", 1, user.GetPaymentMode(model.CREDIT_CARD))
	if err != nil {
		fmt.Println(err)
	}

	user.GetPaymentMode(model.CREDIT_CARD).AddBalance(100)
	err = vendingMachine.BuyProduct("Lays", 1, user.GetPaymentMode(model.CREDIT_CARD))
	if err != nil {
		fmt.Println(err)
	}

}
