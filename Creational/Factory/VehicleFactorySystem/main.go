package main

import (
	"fmt"
	"vehiclefactory/factory"
)

const CAR = "car"
const TRUCK = "truck"

func main() {
	fmt.Println("Vehicle Factory System")

	// Create Factory
	vehicleFactory := factory.NewVehicleFactory()

	// Create Car
	car, err := vehicleFactory.NewVehicle(CAR, "Maruti Suzuki", 4, 5, 1.2)
	if err != nil {
		fmt.Println("There was some error creating car")
	}

	truck, err := vehicleFactory.NewVehicle(TRUCK, "Tata Motors", 4, 5, 2.2)
	if err != nil {
		fmt.Println("There was some error creating truck")
	}

	fmt.Println(car.GetDetails())
	fmt.Println(truck.GetDetails())
}
