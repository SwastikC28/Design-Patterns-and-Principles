### Problem Statement: Vehicle Factory System

Your task is to design a simple Vehicle Factory System that helps in creating different types of vehicles. The system should be able to produce a variety of vehicle types (e.g., Car, Truck, Motorcycle) based on the type requested by the client. Each vehicle should have a method getDetails() that returns a description of the vehicle.

#### Requirements:
Vehicle Interface:

#### Create an interface Vehicle that has a method getDetails() that returns a string description of the vehicle.
Concrete Vehicle Classes:

Implement three concrete classes:
Car with specific details about a car.
Truck with specific details about a truck.
Motorcycle with specific details about a motorcycle.
Factory Class:

Create a VehicleFactory class with a static method createVehicle(String type) that returns an object of type Vehicle. The method should take a string parameter type and return the appropriate Vehicle object (Car, Truck, or Motorcycle).
Client Class:

Design a Client class to demonstrate the use of the factory. The client should call the VehicleFactory to get an instance of the desired Vehicle type and then call getDetails() to display the vehicle details.
Constraints:
You should use the Factory Design Pattern to create a vehicle without specifying the exact class name in the client code.
Ensure the system is scalable and easily extendable if more vehicle types need to be added in the future.
Example:
Input:

Request to create a Car from the factory.
Request to create a Truck from the factory.
Output:

Details of the Car: "This is a Car with 4 doors, 5 seats, and an engine capacity of 2.5L."
Details of the Truck: "This is a Truck with a cargo capacity of 10 tons, 2 seats, and an engine capacity of 5.0L."
Note: Ensure that the client code does not know the concrete class names (Car, Truck, etc.) and only interacts with the Vehicle interface.