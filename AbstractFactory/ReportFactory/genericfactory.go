package main

import (
	product "financefactory/Product"
	"financefactory/factory"
)

type GenericFactory struct{}

func NewGenericFactory(factoryType string) product.ReportFactory {
	switch factoryType {
	case "finance":
		return factory.NewFinanceFactory()
	case "hr":
		return factory.NewHRFactory()
	}

	return nil
}
