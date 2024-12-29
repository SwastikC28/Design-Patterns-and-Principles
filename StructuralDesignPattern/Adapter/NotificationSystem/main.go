package main

import (
	"fmt"
	"noti/adapter"
	"noti/channels"
)

func main() {
	fmt.Println("Notification System")

	// Email notification example
	emailAdapter := adapter.NewEmailNotificationAdapter(&channels.EmailNotification{})
	emailParams := map[string]interface{}{
		"subject":   "Welcome!",
		"body":      "Thank you for signing up.",
		"recipient": "user@example.com",
	}
	emailAdapter.Send(emailParams)

	// SMS notification example
	smsAdapter := adapter.NewSMSNotificationAdapter(&channels.SMSNotification{})
	smsParams := map[string]interface{}{
		"phoneNumber": "1234567890",
		"message":     "Your OTP is 1234.",
	}
	smsAdapter.Send(smsParams)

	// Push notification example
	pushAdapter := adapter.NewPushNotificationAdapter(&channels.PushNotification{})
	pushParams := map[string]interface{}{
		"deviceID": "device123",
		"message":  "You have a new message.",
	}
	pushAdapter.Send(pushParams)
}
