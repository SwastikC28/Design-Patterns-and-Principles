package main

import (
	"carproblem/models"
	"fmt"
)

func main() {

	carBuilder := models.NewCarBuilder().
		SetColor("Red").
		SetEngineManufacturedYear("2023").
		SetEngineVer("v69").
		SetModel("GT").
		SetSunroof(true).
		SetTransmissionType("Manual")

	car := carBuilder.Build()

	fmt.Printf("Your Car is Ready  =   %+v", car)
}
