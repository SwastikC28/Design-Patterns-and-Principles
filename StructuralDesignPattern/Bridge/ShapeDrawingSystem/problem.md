# Bridge Design Pattern Example: Shape Drawing System

## Problem Statement

Design a software system that can draw different shapes (e.g., circles, squares) on different rendering devices (e.g., printers, screens). The system should allow new shapes and rendering devices to be added without modifying the existing code.

## Requirements

1. **Shapes:**
   - Circle
   - Square

2. **Rendering Devices:**
   - Printer
   - Screen

3. **Behavior:**
   - Each shape should have its own way of being drawn.
   - Each rendering device should have its own mechanism for rendering shapes.
   - Ensure that adding new shapes or rendering devices does not require changes to the existing code.

## Constraints

- Use the **Bridge Design Pattern** to separate the abstraction (shapes) from the implementation (rendering devices).
- Focus on extensibility, so that new shapes or rendering devices can be added easily.

## Example Use Case

The **Bridge Pattern** helps decouple the shape abstraction (Circle, Square) from the implementation (Printer, Screen). This design allows flexibility to add new shapes (e.g., Triangle, Rectangle) or devices (e.g., Projector, Web) without affecting the existing system.

## Key Points

- **Shape** is the abstraction.
- **Rendering Device** is the implementor.
- The pattern provides **flexibility and extensibility** by separating the core functionality (shapes) from the platform-specific logic (devices).
