package main

import (
	"fmt"
	"sandwich/model"
)

func main() {
	fmt.Println("Builder Pattern Implemented")
	sandwichBuilder := model.NewSandwichBuilder()

	sandwichBuilder.SetBread("Whole Wheat")
	sandwichBuilder.SetPatty("Veg Paneer Patty")
	sandwichBuilder.SetSauces("mayonaise", "tandoori mayonaise", "red chilly sauce", "mustard sauce")

	// Create Order 1
	orderOne := sandwichBuilder.Build()
	fmt.Printf("Order 1 is ready %v.\n", orderOne)

	sandwichBuilder.SetBread("Honey Bread")
	sandwichBuilder.SetPatty("Chicken Patty")
	sandwichBuilder.SetSauces("tandoori mayonaise", "mayonaise")

	// Create Order 2
	orderTwo := sandwichBuilder.Build()
	fmt.Printf("Order 2 is ready %v.\n", orderTwo)

	fmt.Println("Restaurant finished taking sandwiches order")
}
