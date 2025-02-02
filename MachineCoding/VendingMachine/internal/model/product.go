package model

type Product interface {
	GetName() string
	GetPrice() int
}

type Chip struct {
	name  string
	price int
}

func NewChip(name string, price int) Product {
	return &Chip{
		name:  name,
		price: price,
	}
}

func (c *Chip) GetName() string {
	return c.name
}

func (c *Chip) GetPrice() int {
	return c.price
}

type Drink struct {
	name  string
	price int
}

func NewDrink(name string, price int) Product {
	return &Drink{
		name:  name,
		price: price,
	}
}

func (d *Drink) GetName() string {
	return d.name
}

func (d *Drink) GetPrice() int {
	return d.price
}

type Biscuit struct {
	name  string
	price int
}

func NewBiscuit(name string, price int) Product {
	return &Biscuit{
		name:  name,
		price: price,
	}
}

func (b *Biscuit) GetName() string {
	return b.name
}

func (b *Biscuit) GetPrice() int {
	return b.price
}
