# Problem Statement: Implementing a Payment System with Multiple Gateways

You are tasked with designing a payment processing system for an e-commerce platform. The platform must support multiple third-party payment gateways (e.g., PayPal, Stripe, and Razorpay). However, these gateways have different interfaces and APIs for processing payments.

## Requirements:
1. **Payment Gateways:**
   - Each payment gateway provides a unique method for initiating and processing payments.
   - Example:
     - PayPal: `processPaypalPayment(amount, currency)`
     - Stripe: `makeStripePayment(totalAmount, paymentMethod, customerDetails)`
     - Razorpay: `initiateRazorpayTransaction(amount, customerDetails)`

2. **System Consistency:**
   - The e-commerce platform should interact with all payment gateways using a single interface, `PaymentProcessor`.
   - The interface should expose a method, `processPayment(amount, currency)`.

3. **Implement an Adapter:**
   - Use the Adapter design pattern to wrap each payment gateway with a common implementation of the `PaymentProcessor` interface.
   - The adapter should translate the `processPayment` method call into the appropriate call for the respective payment gateway.

4. **Extendability:**
   - The system should allow easy integration of new payment gateways without modifying existing code.

## Additional Constraints:
- The e-commerce platform should be able to switch between payment gateways dynamically at runtime.
- Ensure the design adheres to the SOLID principles, particularly the Open/Closed Principle.

## Deliverables:
Design the system and demonstrate the integration of at least three payment gateways using the Adapter design pattern.
