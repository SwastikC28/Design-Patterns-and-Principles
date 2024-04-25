package main

import "absfactory/factory"

func main() {
	// Pencil Factory
	pencilFactory := factory.NewAbstractPencilFactory()
	natarajFactory := pencilFactory.GetPencilFactory("nataraj")

	natarajPencil := natarajFactory.GetPencil()
	natarajPencil.Draw()

	// Pen Factory
	penFactory := factory.NewAbstractPenFactory()
	reynoldsFactory := penFactory.GetPenFactory("reynolds")
	reynoldsPen := reynoldsFactory.GetPen()
	reynoldsPen.Sign()
}
