package main

import (
	"fmt"
	"weatherstation/internal/model"
)

func main() {
	fmt.Println("Observer Design Pattern")
	weatherStation := model.NewWeatherStation()

	display1 := model.NewDisplay(1, "LED TV")
	display2 := model.NewDisplay(2, "OLED TV")
	display3 := model.NewDisplay(3, "QLED TV")

	weatherStation.GetNotificationService().Subscribe(display1)
	weatherStation.GetNotificationService().Subscribe(display2)
	weatherStation.GetNotificationService().Subscribe(display3)

	weatherStation.UpdateTemperature(40)
	weatherStation.UpdateHumidity(30)

	weatherStation.GetNotificationService().Unsubscribe(display3)
	weatherStation.UpdatePressure(10)
}
