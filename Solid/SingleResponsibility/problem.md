### Low-Level Design Problem Statement:

Design a class for a simple banking system that adheres to the Single Responsibility Principle (SRP). The class should represent a bank account and be responsible for managing basic operations such as depositing funds, withdrawing funds, and checking the account balance.

The class should have the following responsibilities:

Account Information: Maintain information about the account, including the account holder's name, account number, and current balance. This responsibility includes methods for retrieving and updating this information.

Deposit Management: Implement functionality for depositing funds into the account. This responsibility involves validating the deposit amount and updating the account balance accordingly.

Withdrawal Management: Implement functionality for withdrawing funds from the account. This responsibility involves validating the withdrawal amount, ensuring sufficient funds are available, and updating the account balance accordingly.

Balance Inquiry: Provide a method to check the current balance of the account.

Each method within the class should adhere to a single responsibility, ensuring that the class remains focused and cohesive. Additionally, consider appropriate error handling mechanisms to handle edge cases such as invalid inputs or insufficient funds during withdrawals.

By designing the class with a clear separation of concerns, you will create a more maintainable and extensible banking system that follows the principles of good software design, including the Single Responsibility Principle.