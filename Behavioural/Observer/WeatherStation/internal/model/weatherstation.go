package model

import (
	"fmt"
	"time"
)

type WeatherStation interface {
	NotifyDisplays()
	GetNotificationService() NotificationService
	UpdateHumidity(int)
	UpdateTemperature(int)
	UpdatePressure(int)
}

type WeatherStationImpl struct {
	WeatherData         *Weather
	notificationService NotificationService
}

func NewWeatherStation() WeatherStation {
	return &WeatherStationImpl{
		WeatherData: &Weather{
			Temperature: 20,
			Humidity:    10,
			Pressure:    5,
		},
		notificationService: NewNotificationService(),
	}
}

func (ws *WeatherStationImpl) NotifyDisplays() {
	fmt.Println("Notifying all displays at time ", time.Now().Format("2006-01-02 15:04:05"))
	ws.notificationService.Notify(*ws.WeatherData)
	fmt.Println()
}

func (ws *WeatherStationImpl) GetNotificationService() NotificationService {
	return ws.notificationService
}

func (ws *WeatherStationImpl) UpdateTemperature(temperature int) {
	ws.WeatherData.Temperature = temperature
	ws.NotifyDisplays()
}

func (ws *WeatherStationImpl) UpdateHumidity(humidity int) {
	ws.WeatherData.Humidity = humidity
	ws.NotifyDisplays()
}
func (ws *WeatherStationImpl) UpdatePressure(pressure int) {
	ws.WeatherData.Pressure = pressure
	ws.NotifyDisplays()
}
