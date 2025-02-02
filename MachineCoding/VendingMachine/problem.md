# Vending Machine Design  

## Problem Statement  
Design a **Vending Machine** that allows users to purchase items by inserting money, selecting a product, and receiving change if applicable. The machine should support multiple product types, maintain an inventory, and handle different denominations of money. The system should be scalable and allow for future enhancements such as digital payments, discount coupons, or loyalty programs.  

## Functional Requirements  
### 1. Display and Selection  
- Show available products along with their prices.  
- Allow users to select a product.  

### 2. Payment Handling  
- Accept different denominations of currency (coins and bills).  
- Validate the inserted amount against the selected product price.  
- Provide change if needed.  

### 3. Inventory Management  
- Track the stock of each product.  
- Prevent the selection of out-of-stock products.  
- Allow an admin to restock items.  

### 4. Transaction Handling  
- Ensure a successful transaction deducts the inventory.  
- If insufficient funds are inserted, allow the user to cancel and get a refund.  

### 5. Error Handling  
- Handle scenarios such as product unavailability, insufficient funds, or machine malfunctions.  

## Non-Functional Requirements  
- The system should be **maintainable and extendable** for new features.  
- The vending machine should have a **simple and intuitive interface**.  
- The machine should provide **fast and reliable** transactions.  

## Future Considerations (Not mandatory for MVP but good to discuss)  
- Support for **digital payments** (credit/debit cards, mobile wallets).  
- **User authentication** for personalized recommendations or loyalty rewards.  
- **Remote monitoring** for tracking inventory and machine health.  
