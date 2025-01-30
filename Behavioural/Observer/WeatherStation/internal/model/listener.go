package model

import (
	"fmt"
	"math"
)

type Observer interface {
	GetID() int
	Update(interface{})
}

type Display struct {
	ID       int
	Name     string
	MaxTemp  int
	MinTemp  int
	Sum      int
	Readings int
}

func NewDisplay(id int, name string) Observer {
	return &Display{
		ID:       id,
		Name:     name,
		MaxTemp:  math.MinInt,
		MinTemp:  math.MaxInt,
		Sum:      0,
		Readings: 0,
	}
}

func (d *Display) GetID() int {
	return d.ID
}

func (d *Display) Update(weather interface{}) {
	fmt.Println("CURRENT WEATHER FOR DISPLAY ", d.ID)

	temperature := weather.(Weather).Temperature

	fmt.Println("TEMPERATURE ", temperature)
	fmt.Println("HUMIDITY ", weather.(Weather).Humidity)
	fmt.Println("PRESSURE ", weather.(Weather).Pressure)

	d.MaxTemp = max(d.MaxTemp, temperature)
	d.MinTemp = min(d.MinTemp, temperature)

	d.Sum = d.Sum + temperature
	d.Readings = d.Readings + 1
	var average float64 = float64(d.Sum) / float64(d.Readings)

	fmt.Println("STATISTICAL INFORMATION FOR DISPLAY ", d.ID)
	fmt.Println("MAX TEMP", d.MaxTemp)
	fmt.Println("MIN TEMP", d.MinTemp)
	fmt.Println("AVG TEMP", average)

}
