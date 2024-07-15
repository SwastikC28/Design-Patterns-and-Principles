package product

type Report interface {
	GenerateReport()
}

type ReportFactory interface {
	NewPDFReport() Report
	NewExcelReport() Report
	NewHTMLReport() Report
}
