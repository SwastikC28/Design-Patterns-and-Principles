package factory

import (
	product "financefactory/Product"
	"fmt"
)

type HRFactory struct{}

func NewHRFactory() *HRFactory {
	return &HRFactory{}
}

func (factory *HRFactory) NewPDFReport() product.Report {
	return &HRPDFReport{}
}

func (factory *HRFactory) NewExcelReport() product.Report {
	return &HRExcelReport{}
}

func (factory *HRFactory) NewHTMLReport() product.Report {
	return &HRHTMLReport{}
}

// Implementation of HR Reports

type HRPDFReport struct{}

func (report *HRPDFReport) GenerateReport() {
	fmt.Println("Generating Report for HR Factory in PDF")
}

type HRExcelReport struct{}

func (report *HRExcelReport) GenerateReport() {
	fmt.Println("Generating Report for HR Factory in Excel")
}

type HRHTMLReport struct{}

func (report *HRHTMLReport) GenerateReport() {
	fmt.Println("Generating Report for HR Factory in HTML")
}
