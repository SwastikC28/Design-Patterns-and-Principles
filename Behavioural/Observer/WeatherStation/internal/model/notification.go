package model

import "fmt"

type NotificationService interface {
	Subscribe(Observer)
	Unsubscribe(Observer)
	Notify(Weather)
}

type NotificationServiceImpl struct {
	observers map[int]Observer
}

func NewNotificationService() NotificationService {
	return &NotificationServiceImpl{
		observers: make(map[int]Observer),
	}
}

func (n *NotificationServiceImpl) Subscribe(display Observer) {
	n.observers[display.GetID()] = display
}

func (n *NotificationServiceImpl) Unsubscribe(display Observer) {
	delete(n.observers, display.GetID())
}

func (n *NotificationServiceImpl) Notify(weather Weather) {
	for _, observer := range n.observers {
		fmt.Println()
		observer.Update(weather)
		fmt.Println()
	}
}
