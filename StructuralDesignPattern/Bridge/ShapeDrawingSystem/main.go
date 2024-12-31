package main

import (
	"fmt"
	"shapedrawingsystem/abstraction"
	"shapedrawingsystem/implementation"
)

func main() {
	fmt.Println("Shape Drawing System")

	// Create platform objects
	windowsPlatform := implementation.NewWindowsPlatform()
	macOSPlatform := implementation.NewMacOSPlatform()
	linuxPlatform := implementation.NewLinuxPlatform()

	// Create and draw shapes on different platforms
	circle1 := abstraction.NewCircle(windowsPlatform)
	circle1.DrawShape()

	circle2 := abstraction.NewCircle(macOSPlatform)
	circle2.DrawShape()

	circle3 := abstraction.NewCircle(linuxPlatform)
	circle3.DrawShape()

	// Optionally, add square shapes as well
	square1 := abstraction.NewSquare(windowsPlatform)
	square1.DrawShape()

	square2 := abstraction.NewSquare(macOSPlatform)
	square2.DrawShape()

	square3 := abstraction.NewSquare(linuxPlatform)
	square3.DrawShape()
}
