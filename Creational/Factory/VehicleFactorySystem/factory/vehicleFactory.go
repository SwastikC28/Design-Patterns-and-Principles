package factory

import (
	"errors"
	"vehiclefactory/models"
)

type VehicleFactory struct {
}

func NewVehicleFactory() *VehicleFactory {
	return &VehicleFactory{}
}

func (vf *VehicleFactory) NewVehicle(vehicle string, name string, door int, seats int, engineCapacity float64) (models.Vehicle, error) {
	switch vehicle {
	case "car":
		return models.NewCar(name, door, seats, engineCapacity), nil
	case "truck":
		return models.NewTruck(name, door, seats, engineCapacity), nil
	default:
		return nil, errors.New("unknown vehicle type")
	}
}
