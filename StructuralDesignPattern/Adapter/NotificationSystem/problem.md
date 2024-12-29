# Problem Statement: Implementing a Notification System

You are tasked with designing a notification system for a mobile app that can send notifications via multiple channels, such as email, SMS, and push notifications. Each notification channel has a different way of sending messages and may have different API requirements.

## Requirements:
1. **Notification Channels:**
   - Each notification channel has its own method of sending notifications.
   - Example:
     - Email: `sendEmail(subject, body, recipient)`
     - SMS: `sendSMS(phoneNumber, message)`
     - Push Notification: `sendPushNotification(deviceID, message)`

2. **System Consistency:**
   - The notification system should provide a unified interface, `NotificationSender`, that allows sending notifications through any channel using a method `sendNotification(message string)`.

3. **Implement an Adapter:**
   - Use the Adapter design pattern to create an adapter for each notification channel that conforms to the `NotificationSender` interface.
   - Each adapter should translate the `sendNotification` method call to the corresponding method of the specific notification channel.

4. **Extendability:**
   - The system should allow easy integration of new notification channels without modifying existing code. For instance, you should be able to add a "Slack" notification channel in the future without changing the current classes.

## Additional Constraints:
- The system should be able to send notifications through any of the available channels dynamically at runtime.
- Ensure the design adheres to the SOLID principles, particularly the Open/Closed Principle.

## Deliverables:
Design the system and implement adapters for at least three notification channels: Email, SMS, and Push Notification. Demonstrate the integration of these adapters using the `NotificationSender` interface.
