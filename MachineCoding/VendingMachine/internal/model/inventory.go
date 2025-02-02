package model

import "fmt"

type Inventory interface {
	DisplayProducts()
	GetProduct(product string) (Product, error)
	GetStock(product string) (int, error)
	Restock(user *User, product Product, quantity int) error
	DeductStock(product Product, quantity int) error
}

type Stock struct {
	product  Product
	quantity int
}

func NewInventory() Inventory {
	return &VendingMachineInventory{
		stock: make(map[string]*Stock),
	}
}

type VendingMachineInventory struct {
	stock map[string]*Stock
}

func (inventory *VendingMachineInventory) GetProduct(product string) (Product, error) {
	p, exists := inventory.stock[product]
	if !exists {
		return nil, fmt.Errorf("no product with that name")
	}

	return p.product, nil
}

func (inventory *VendingMachineInventory) GetStock(product string) (int, error) {
	p, exists := inventory.stock[product]
	if !exists {
		return -1, fmt.Errorf("no product with that name")
	}

	return p.quantity, nil
}

func (inventory *VendingMachineInventory) Restock(user *User, product Product, quantity int) error {
	if !user.isAdmin {
		return fmt.Errorf("user unauthorized")
	}

	if quantity < 0 {
		return fmt.Errorf("invalid quantity")
	}

	stock, exists := inventory.stock[product.GetName()]
	if !exists {
		inventory.stock[product.GetName()] = &Stock{
			product:  product,
			quantity: quantity,
		}

		return nil
	}

	stock.quantity += quantity
	return nil
}

func (inventory *VendingMachineInventory) DeductStock(product Product, quantity int) error {
	p, exists := inventory.stock[product.GetName()]
	if !exists {
		return fmt.Errorf("no product with that name")
	}

	if p.quantity < quantity {
		return fmt.Errorf("couldn't process due to low stock")
	}

	inventory.stock[product.GetName()].quantity -= quantity
	return nil
}

func (inventory *VendingMachineInventory) DisplayProducts() {
	for product, stock := range inventory.stock {
		fmt.Println("Product", product, fmt.Sprintf("₹%d", stock.product.GetPrice()), "Q -", stock.quantity)
	}
}
