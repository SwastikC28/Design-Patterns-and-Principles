package model

import "errors"

type TransactionHandler interface {
	Handle(productName string, quantity int, paymentMode PaymentMode) error
	DeductStock(Product, int) error
}

type transactionHandler struct {
	paymentHandler PaymentHandler
	inventory      Inventory
	bank           VendingMachineBank
}

func NewTransactionHandler(inventory Inventory) *transactionHandler {
	return &transactionHandler{
		inventory: inventory,
	}
}

func (t *transactionHandler) GetStock(product string) (int, error) {
	stock, err := t.inventory.GetStock(product)
	if err != nil {
		return -1, err
	}

	return stock, nil
}

func (t *transactionHandler) DeductStock(product Product, quantity int) error {
	stock, err := t.inventory.GetStock(product.GetName())
	if err != nil {
		return err
	}

	if stock < quantity {
		return errors.New("low stock")
	}

	return t.inventory.DeductStock(product, quantity)
}

func (t *transactionHandler) Handle(productName string, quantity int, paymentMode PaymentMode) error {

	product, _ := t.inventory.GetProduct(productName)

	amount := quantity * product.GetPrice()

	// Get payment
	err := t.paymentHandler.Pay(paymentMode, amount)
	if err != nil {
		return err
	}

	// Deduct stock
	err = t.inventory.DeductStock(product, quantity)
	if err != nil {
		t.paymentHandler.Refund(paymentMode, amount)
		return err
	}

	t.bank.AddBalance(amount)
	return nil
}
