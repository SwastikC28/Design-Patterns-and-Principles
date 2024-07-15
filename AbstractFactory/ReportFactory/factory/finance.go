package factory

import (
	product "financefactory/Product"
	"fmt"
)

type FinanceFactory struct{}

func NewFinanceFactory() *FinanceFactory {
	return &FinanceFactory{}
}

func (factory *FinanceFactory) NewPDFReport() product.Report {
	return &FinancePDFReport{}
}

func (factory *FinanceFactory) NewExcelReport() product.Report {
	return &FinanceExcelReport{}
}

func (factory *FinanceFactory) NewHTMLReport() product.Report {
	return &FinanceHTMLReport{}
}

// Implementation of Finance Reports

type FinancePDFReport struct{}

func (report *FinancePDFReport) GenerateReport() {
	fmt.Println("Generating Report for Finance Factory in PDF")
}

type FinanceExcelReport struct{}

func (report *FinanceExcelReport) GenerateReport() {
	fmt.Println("Generating Report for Finance Factory in Excel")
}

type FinanceHTMLReport struct{}

func (report *FinanceHTMLReport) GenerateReport() {
	fmt.Println("Generating Report for Finance Factory in HTML")
}
