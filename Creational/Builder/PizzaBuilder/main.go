package main

import (
	"fmt"
	"pizza-builder/model"
)

func main() {
	fmt.Println("Pizza Builder Problem")

	pizzaBuilder := model.NewPizzaBuilder()

	// Example: Valid pizza configuration
	pizza := pizzaBuilder.
		Base("thin crust").
		Cheese("mozzarella").
		Sauce("tomato sauce").
		Size("medium").
		Toppings("onion, bell peppers").
		Build()

	fmt.Printf("Your order of pizza is completed: %+v\n", pizza)

	// Example: Invalid values (defaults will be applied)
	invalidPizza := pizzaBuilder.
		Base("random crust").
		Cheese("random cheese").
		Sauce("random sauce").
		Size("extra-large").
		Toppings("random toppings").
		Build()

	fmt.Printf("Your order with invalid inputs: %+v\n", invalidPizza)
}
