package shapefactory

import "factory-creational/shape"

type ShapeFactory interface {
	GetShape(shapeType string) shape.Shape
}

type ShapeFactoryImpl struct{}

func NewShapeFactory() ShapeFactory {
	return &ShapeFactoryImpl{}
}

func (s *ShapeFactoryImpl) GetShape(shapeType string) shape.Shape {
	if shapeType == "square" {
		return &shape.Square{}
	} else if shapeType == "rectangle" {
		return &shape.Rectangle{}
	}

	return nil
}
