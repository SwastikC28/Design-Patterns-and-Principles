package main

import (
	shapefactory "factory-creational/shapeFactory"
	"fmt"
)

func main() {
	var ShapeFactory = shapefactory.NewShapeFactory()

	rec := ShapeFactory.GetShape("rectangle")
	square := ShapeFactory.GetShape("square")

	fmt.Printf("%+v %+v", rec, square)
}
