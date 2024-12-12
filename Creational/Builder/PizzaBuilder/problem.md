# Problem Statement: Implement a Pizza Builder

You are tasked with creating a system to order pizzas in a customizable way. Each pizza consists of the following components:

1. **Base**: Options include thin crust, thick crust, or cheese crust.  
2. **Sauce**: Options include tomato sauce, white sauce, or pesto sauce.  
3. **Toppings**: Customers can choose any combination of toppings, such as pepperoni, mushrooms, olives, jalapeños, onions, and bell peppers.  
4. **Cheese**: Options include mozzarella, cheddar, or no cheese.  
5. **Size**: The pizza can be small, medium, or large.

## Requirements
- The pizza should only be created after the customer finalizes their order.
- Each pizza should include all the components mentioned above, but customers can omit any category (e.g., no sauce, no toppings).
- The system should ensure that each pizza is immutable once built.
- Write a program to design and build multiple pizzas efficiently using the builder design pattern.

### Bonus Challenge
Add support for a "combo deal," where customers can order a predefined set of pizzas with common configurations.
