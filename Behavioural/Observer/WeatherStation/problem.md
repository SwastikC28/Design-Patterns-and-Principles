# Observer Pattern Practice Problem

## Problem Statement

You need to implement a **Weather Station** system using the **Observer Pattern**. The system should have the following features:

### 1. Weather Station (Subject)
- Maintains temperature, humidity, and pressure readings.  
- Notifies all registered displays whenever weather data changes.  

### 2. Weather Displays (Observers)
- Multiple displays can subscribe to the Weather Station.  
- Each display updates and shows the latest weather conditions when notified.  
- Implement at least three types of displays:  
  - **Current Conditions Display:** Shows the latest temperature, humidity, and pressure.  
  - **Statistics Display:** Tracks and displays the average, min, and max temperature.  

### 3. Dynamic Subscription Management
- Displays should be able to subscribe (attach) and unsubscribe (detach) from the Weather Station dynamically.

### 4. Push or Pull Mechanism
- The Weather Station should notify observers either by pushing updated data or allowing them to pull the latest values.

### 5. Driver Code
- Simulate the system by creating a Weather Station and registering different displays.
- Update weather conditions multiple times and verify if all displays update correctly.
