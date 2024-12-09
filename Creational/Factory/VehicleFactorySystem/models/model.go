package models

import "fmt"

type Vehicle interface {
	GetDetails() string
}

type Car struct {
	name           string
	door           int
	seats          int
	engineCapacity float64
}

func NewCar(name string, door int, seats int, engineCapacity float64) Vehicle {
	return &Car{
		name:           name,
		door:           door,
		seats:          seats,
		engineCapacity: engineCapacity,
	}
}

func (c *Car) GetDetails() string {
	return fmt.Sprintf("This car is %s. It has %d doors, %d seats and the engine capacity is %f", c.name, c.door, c.seats, c.engineCapacity)
}

// -----------------------------------------------------

type Truck struct {
	name           string
	door           int
	seats          int
	engineCapacity float64
}

func NewTruck(name string, door int, seats int, engineCapacity float64) Vehicle {
	return &Truck{
		name:           name,
		door:           door,
		seats:          seats,
		engineCapacity: engineCapacity,
	}
}

func (c *Truck) GetDetails() string {
	return fmt.Sprintf("This truck is %s. It has %d doors, %d seats and the engine capacity is %f", c.name, c.door, c.seats, c.engineCapacity)
}

// ------------------------------------------------------
