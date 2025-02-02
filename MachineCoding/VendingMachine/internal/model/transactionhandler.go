package model

import "errors"

type TransactionHandler interface {
	DeductStock(Product, int) error
}

type transactionHandler struct {
	inventory Inventory
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
