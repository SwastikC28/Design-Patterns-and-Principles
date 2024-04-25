package factory

import "absfactory/product"

type AbstractPencilFactory struct{}

func NewAbstractPencilFactory() *AbstractPencilFactory {
	return &AbstractPencilFactory{}
}

func (ap *AbstractPencilFactory) GetPencilFactory(pencilType string) PencilFactory {
	if pencilType == "apsara" {
		return &ApsaraPencilFactory{}
	} else {
		return &NatarajPencilFactory{}
	}
}

type PencilFactory interface {
	GetPencil() product.Pencil
}

type ApsaraPencilFactory struct{}

func (ap *ApsaraPencilFactory) GetPencil() product.Pencil {
	return &product.Apsara{}
}

type NatarajPencilFactory struct{}

func (n *NatarajPencilFactory) GetPencil() product.Pencil {
	return &product.Nataraj{}
}
