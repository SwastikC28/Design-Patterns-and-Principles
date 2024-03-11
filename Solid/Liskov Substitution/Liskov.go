package main

type Shape interface {
	perimeter() float64
	area() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) perimeter() float64 {
	return c.radius * 2 * 22 * 7
}

func (c *Circle) area() float64 {
	return 22 * 7 * c.radius * c.radius
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

type Square struct {
	side float64
}

func (s *Square) perimeter() float64 {
	return 4 * s.side
}

func (s *Square) area() float64 {
	return s.side * s.side
}

func NewSquare(side float64) *Square {
	return &Square{side: side}
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.breadth + r.length)
}

func (r *Rectangle) area() float64 {
	return r.length * r.breadth
}

func NewRectangle(length, breadth float64) *Rectangle {
	return &Rectangle{
		length:  length,
		breadth: breadth,
	}
}

func main() {
	var rectangle Shape = NewRectangle(2, 3)
	var square Shape = NewSquare(2)
	var circle Shape = NewCircle(7)

	shapes := []Shape{rectangle, square, circle}

	maxArea := 0.0
	maxPerimeter := 0.0

	for _, shape := range shapes {
		area := shape.area()
		peri := shape.perimeter()

		maxArea = max(maxArea, area)
		maxPerimeter = max(peri, maxPerimeter)
	}

}
