package factory

import "absfactory/product"

type AbstractPenFactory struct{}

func NewAbstractPenFactory() *AbstractPenFactory {
	return &AbstractPenFactory{}
}

func (ap *AbstractPenFactory) GetPenFactory(penType string) PenFactory {
	if penType == "reynolds" {
		return &ReynoldsPenFactory{}
	} else {
		return &PilotPenFactory{}
	}
}

type PenFactory interface {
	GetPen() product.Pen
}

type ReynoldsPenFactory struct{}

func (r *ReynoldsPenFactory) GetPen() product.Pen {
	return &product.Reynolds{}
}

type PilotPenFactory struct{}

func (n *PilotPenFactory) GetPen() product.Pen {
	return &product.Pilot{}
}
