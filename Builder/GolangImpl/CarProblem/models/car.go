package models

type Car struct {
	model                  string
	color                  string
	engineVer              string
	engineManufacturedYear string
	transmissionType       string
	sunroof                bool
}

type CarBuilder interface {
	SetModel(string) CarBuilder
	SetColor(string) CarBuilder
	SetEngineVer(string) CarBuilder
	SetEngineManufacturedYear(string) CarBuilder
	SetTransmissionType(string) CarBuilder
	SetSunroof(bool) CarBuilder
	Build() *Car
}

type CarBuilderImpl struct {
	Model                  string
	Color                  string
	EngineVer              string
	EngineManufacturedYear string
	TransmissionType       string
	Sunroof                bool
}

func NewCarBuilder() CarBuilder {
	return &CarBuilderImpl{}
}

func (builder *CarBuilderImpl) SetModel(model string) CarBuilder {
	builder.Model = model
	return builder
}

func (builder *CarBuilderImpl) SetColor(color string) CarBuilder {
	builder.Color = color
	return builder
}

func (builder *CarBuilderImpl) SetEngineVer(engineVersion string) CarBuilder {
	builder.EngineVer = engineVersion
	return builder
}

func (builder *CarBuilderImpl) SetEngineManufacturedYear(yom string) CarBuilder {
	builder.EngineManufacturedYear = yom
	return builder
}

func (builder *CarBuilderImpl) SetTransmissionType(transmissionType string) CarBuilder {
	builder.TransmissionType = transmissionType
	return builder
}

func (builder *CarBuilderImpl) SetSunroof(hasSunroof bool) CarBuilder {
	builder.Sunroof = hasSunroof
	return builder
}

func (builder *CarBuilderImpl) Build() *Car {
	car := &Car{
		model:                  builder.Model,
		color:                  builder.Color,
		engineVer:              builder.EngineVer,
		engineManufacturedYear: builder.EngineManufacturedYear,
		transmissionType:       builder.TransmissionType,
		sunroof:                builder.Sunroof,
	}

	return car
}
