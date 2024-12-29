package channels

import "fmt"

type EmailNotification struct{}

func (notification *EmailNotification) Send(subject, body, recipient string) error {
	fmt.Printf("Email sent to %s with subject '%s' and body '%s'\n", recipient, subject, body)
	return nil
}

type SMSNotification struct{}

func (notification *SMSNotification) Send(phoneNumber, message string) error {
	fmt.Printf("SMS sent to %s with message '%s'\n", phoneNumber, message)
	return nil
}

type PushNotification struct{}

func (notification *PushNotification) Send(deviceID, message string) error {
	fmt.Printf("Push notification sent to device %s with message '%s'\n", deviceID, message)
	return nil
}
