package model

type Pizza struct {
	base     string
	sauce    string
	toppings string
	cheese   string
	size     string
}

type PizzaBuilder struct {
	base     string
	sauce    string
	toppings string
	cheese   string
	size     string
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

func (pb *PizzaBuilder) Base(base string) *PizzaBuilder {
	validBases := map[string]bool{"thin crust": true, "cheese crust": true, "thick crust": true}
	if !validBases[base] {
		base = "thin crust" // Default Base
	}
	pb.base = base
	return pb
}

func (pb *PizzaBuilder) Sauce(sauce string) *PizzaBuilder {
	validSauces := map[string]bool{"tomato sauce": true, "white sauce": true, "pesto sauce": true}
	if !validSauces[sauce] {
		sauce = "tomato sauce" // Default Sauce
	}
	pb.sauce = sauce
	return pb
}

func (pb *PizzaBuilder) Toppings(toppings string) *PizzaBuilder {
	// Add additional validation if needed
	pb.toppings = toppings
	return pb
}

func (pb *PizzaBuilder) Cheese(cheese string) *PizzaBuilder {
	validCheeses := map[string]bool{"mozzarella": true, "cheddar": true, "no cheese": true}
	if !validCheeses[cheese] {
		cheese = "no cheese" // Default Cheese
	}
	pb.cheese = cheese
	return pb
}

func (pb *PizzaBuilder) Size(size string) *PizzaBuilder {
	validSizes := map[string]bool{"small": true, "medium": true, "large": true}
	if !validSizes[size] {
		size = "small" // Default Size
	}
	pb.size = size
	return pb
}

func (pb *PizzaBuilder) Build() *Pizza {
	pizza := &Pizza{
		base:     pb.base,
		sauce:    pb.sauce,
		toppings: pb.toppings,
		cheese:   pb.cheese,
		size:     pb.size,
	}

	// Reset builder's state to ensure no leftover configuration
	pb.base = ""
	pb.sauce = ""
	pb.toppings = ""
	pb.cheese = ""
	pb.size = ""

	return pizza

}
