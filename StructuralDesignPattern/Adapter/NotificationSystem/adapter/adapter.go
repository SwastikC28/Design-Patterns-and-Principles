package adapter

import "noti/channels"

// NotificationSystem is the unified interface.
type NotificationSystem interface {
	Send(params map[string]interface{}) error
}

// EmailNotificationAdapter adapts EmailNotification to NotificationSystem.
type EmailNotificationAdapter struct {
	emailNotification *channels.EmailNotification
}

func NewEmailNotificationAdapter(emailNotification *channels.EmailNotification) NotificationSystem {
	return &EmailNotificationAdapter{emailNotification: emailNotification}
}

func (adapter *EmailNotificationAdapter) Send(params map[string]interface{}) error {
	subject := params["subject"].(string)
	body := params["body"].(string)
	recipient := params["recipient"].(string)
	return adapter.emailNotification.Send(subject, body, recipient)
}

// SMSNotificationAdapter adapts SMSNotification to NotificationSystem.
type SMSNotificationAdapter struct {
	smsNotification *channels.SMSNotification
}

func NewSMSNotificationAdapter(smsNotification *channels.SMSNotification) NotificationSystem {
	return &SMSNotificationAdapter{smsNotification: smsNotification}
}

func (adapter *SMSNotificationAdapter) Send(params map[string]interface{}) error {
	phoneNumber := params["phoneNumber"].(string)
	message := params["message"].(string)
	return adapter.smsNotification.Send(phoneNumber, message)
}

// PushNotificationAdapter adapts PushNotification to NotificationSystem.
type PushNotificationAdapter struct {
	pushNotification *channels.PushNotification
}

func NewPushNotificationAdapter(pushNotification *channels.PushNotification) NotificationSystem {
	return &PushNotificationAdapter{pushNotification: pushNotification}
}

func (adapter *PushNotificationAdapter) Send(params map[string]interface{}) error {
	deviceID := params["deviceID"].(string)
	message := params["message"].(string)
	return adapter.pushNotification.Send(deviceID, message)
}
