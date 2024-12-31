package abstraction

import "shapedrawingsystem/implementation"

type Shape interface {
	DrawShape()
}

// ==============================================================
type Circle struct {
	platform implementation.Platform
}

func NewCircle(platform implementation.Platform) Shape {
	return &Circle{
		platform: platform,
	}
}

func (shape *Circle) DrawShape() {
	shape.platform.RenderShape("Circle")
}

// ==============================================================
type Square struct {
	platform implementation.Platform
}

func NewSquare(platform implementation.Platform) Shape {
	return &Square{
		platform: platform,
	}
}

func (shape *Square) DrawShape() {
	shape.platform.RenderShape("Square")

}
