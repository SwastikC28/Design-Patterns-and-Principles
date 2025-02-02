package model

import "fmt"

type VendingMachine interface {
	BuyProduct(product string, quantity int, paymentMode PaymentMode) error
	DisplayProducts()
	AddProduct(product Product, quantity int)
}

type vendingMachine struct {
	inventory          Inventory
	paymentHandler     PaymentHandler
	transactionHandler TransactionHandler
}

func NewVendingMachine(inventory Inventory, bank VendingMachineBank) VendingMachine {
	return &vendingMachine{
		inventory:          inventory,
		paymentHandler:     NewPaymentHandler(bank),
		transactionHandler: NewTransactionHandler(inventory),
	}
}

func (machine *vendingMachine) BuyProduct(product string, quantity int, paymentMode PaymentMode) error {
	inventoryProduct, err := machine.inventory.GetProduct(product)
	if err != nil {
		return err
	}

	stock, err := machine.inventory.GetStock(product)
	if err != nil {
		return err
	}

	if stock < quantity {
		return fmt.Errorf("low stock")
	}

	err = machine.paymentHandler.Pay(paymentMode, quantity)
	if err != nil {
		return err
	}

	err = machine.transactionHandler.DeductStock(inventoryProduct, quantity)
	if err != nil {
		machine.paymentHandler.Refund(paymentMode, quantity)
		return err
	}

	fmt.Println("Product dropped.")
	return nil
}

func (machine *vendingMachine) DisplayProducts() {
	machine.inventory.DisplayProducts()
}

func (machine *vendingMachine) AddProduct(product Product, quantity int) {}
