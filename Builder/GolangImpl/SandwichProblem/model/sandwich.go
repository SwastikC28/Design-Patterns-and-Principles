package model

type Sandwich struct {
	Patty  string
	Bread  string
	Sauces []string
}

type SandwichBuilder interface {
	SetPatty(string) SandwichBuilder
	SetBread(string) SandwichBuilder
	SetSauces(...string) SandwichBuilder
	Build() *Sandwich
}

type SandwichBuilderImpl struct {
	Patty  string
	Bread  string
	Sauces []string
}

func NewSandwichBuilder() SandwichBuilder {
	return &SandwichBuilderImpl{}
}

func (builder *SandwichBuilderImpl) SetPatty(patty string) SandwichBuilder {
	builder.Patty = patty
	return builder
}

func (builder *SandwichBuilderImpl) SetBread(bread string) SandwichBuilder {
	builder.Bread = bread
	return builder
}

func (builder *SandwichBuilderImpl) SetSauces(Sauces ...string) SandwichBuilder {
	builder.Sauces = Sauces
	return builder
}

func (builder *SandwichBuilderImpl) Build() *Sandwich {

	// Set default values for mandatory fields if not provided
	if builder.Patty == "" {
		builder.Patty = "Veg Patty"
	}
	if builder.Bread == "" {
		builder.Bread = "Whole wheat Bread"
	}

	sandwich := &Sandwich{
		Sauces: builder.Sauces,
		Bread:  builder.Bread,
		Patty:  builder.Patty,
	}
	return sandwich
}
