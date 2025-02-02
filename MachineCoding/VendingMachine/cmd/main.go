package main

import (
	"fmt"
	"os"
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

	for {
		fmt.Println("\nAvailable actions:")
		fmt.Println("1. Display Products")
		fmt.Println("2. Buy Product")
		fmt.Println("3. Add Credit to Card")
		fmt.Println("4. Show Sales")
		fmt.Println("5. Exit")

		fmt.Print("\nChoose an option: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			vendingMachine.DisplayProducts()

		case 2:
			fmt.Print("Enter the product name: ")
			var productName string
			fmt.Scan(&productName)

			fmt.Print("Enter the quantity: ")
			var quantity int
			fmt.Scan(&quantity)

			fmt.Println("Choose payment method:")
			fmt.Println("1. Credit Card")
			fmt.Println("2. Cash")
			fmt.Print("Enter payment method (1/2): ")
			var paymentChoice int
			fmt.Scan(&paymentChoice)

			var paymentMethod model.PaymentMode
			if paymentChoice == 1 {
				paymentMethod = user.GetPaymentMode(model.CREDIT_CARD)
			} else {
				paymentMethod = user.GetPaymentMode(model.CASH)
			}

			err := vendingMachine.BuyProduct(productName, quantity, paymentMethod)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 3:
			fmt.Print("Enter amount to add to credit card: ")
			var amount float64
			fmt.Scan(&amount)

			user.GetPaymentMode(model.CREDIT_CARD).AddBalance(int(amount))
			fmt.Println("Amount added to credit card.")

		case 4:
			vendingMachine.ShowVendingMachineSales()

		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
