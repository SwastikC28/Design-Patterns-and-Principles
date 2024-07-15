package main

import "fmt"

func main() {
	fmt.Println("Report Generation Factory")

	financeReportFactory := NewGenericFactory("finance")
	financeReportFactory.NewExcelReport().GenerateReport()
	financeReportFactory.NewHTMLReport().GenerateReport()
	financeReportFactory.NewPDFReport().GenerateReport()

	hrReportFactory := NewGenericFactory("hr")
	hrReportFactory.NewExcelReport().GenerateReport()
	hrReportFactory.NewHTMLReport().GenerateReport()
	hrReportFactory.NewPDFReport().GenerateReport()

}
