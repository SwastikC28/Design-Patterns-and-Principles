package shape

type Shape interface {
	getPerimeter() int
	getArea() int
}

type Rectangle struct {
	length  int
	breadth int
}

func (r *Rectangle) getPerimeter() int {
	return (r.length + r.breadth) * 2
}

func (r *Rectangle) getArea() int {
	return r.length * r.breadth
}

type Square struct {
	side int
}

func (s *Square) getPerimeter() int {
	return s.side * 4
}

func (s *Square) getArea() int {
	return s.side * s.side
}
