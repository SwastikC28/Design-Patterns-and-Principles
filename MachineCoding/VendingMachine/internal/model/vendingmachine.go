package model

import "fmt"

type VendingMachine interface {
	BuyProduct(product string, quantity int, paymentMode PaymentMode) error
	DisplayProducts()
	ShowVendingMachineSales()
}

type vendingMachine struct {
	inventory          Inventory
	paymentHandler     PaymentHandler
	transactionHandler TransactionHandler
	bank               VendingMachineBank
}

func NewVendingMachine(inventory Inventory, bank VendingMachineBank) VendingMachine {
	return &vendingMachine{
		inventory:          inventory,
		paymentHandler:     NewPaymentHandler(bank),
		transactionHandler: NewTransactionHandler(inventory),
		bank:               bank,
	}
}

func (machine *vendingMachine) BuyProduct(product string, quantity int, paymentMode PaymentMode) error {
	stock, err := machine.inventory.GetStock(product)
	if err != nil {
		return err
	}

	if stock < quantity {
		return fmt.Errorf("low stock")
	}

	err = machine.transactionHandler.Handle(product, quantity, paymentMode)
	if err != nil {
		return nil
	}

	fmt.Println("Product dropped.")
	return nil
}

func (machine *vendingMachine) DisplayProducts() {
	machine.inventory.DisplayProducts()
}

func (machine *vendingMachine) ShowVendingMachineSales() {
	fmt.Println("Machine Sales", machine.bank.GetBalance())
}
