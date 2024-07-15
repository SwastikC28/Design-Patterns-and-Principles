Create a content management system (CMS) that supports generating different types of reports (e.g., PDF reports, Excel reports, HTML reports) for various departments (e.g., Finance, HR, Sales). Use the Abstract Factory design pattern to ensure that the appropriate report generators are created for each department, maintaining consistency and scalability.

Requirements:
Abstract Factory Interface:

Define an interface (ReportFactory) for creating abstract products (PDFReport, ExcelReport, HTMLReport).
Concrete Factories:

Implement concrete factory classes for each department (FinanceReportFactory, HRReportFactory, SalesReportFactory) that implement the ReportFactory interface. Each factory should create report generators specific to its department.
Abstract Product Interfaces:

Define interfaces for the products (PDFReport, ExcelReport, HTMLReport) that the concrete factories will produce.
Concrete Products:

Implement concrete product classes for each report type and department. For example, FinancePDFReport, HRPDFReport, SalesPDFReport, FinanceExcelReport, HRExcelReport, SalesExcelReport, FinanceHTMLReport, HRHTMLReport, SalesHTMLReport.
Client Code:

Create client code that uses the ReportFactory interface to create reports. The client code should not depend on the concrete classes but should work with any factory that implements the ReportFactory interface.
Steps:
Define the Abstract Factory Interface:

Create an interface for the report factory that declares methods for creating the different types of reports.
Define Abstract Product Interfaces:

Create interfaces for the various types of reports (PDF, Excel, HTML) with methods related to report generation.
Implement Concrete Factories:

Implement factory classes for each department that produce the corresponding report objects.
Implement Concrete Products:

Implement classes for the concrete report types, ensuring that each department has its specific implementations.
Client Code:

Write client code to demonstrate the usage of the abstract factory pattern to generate reports for different departments and types. The client should be able to request a report from any department and in any format without knowing the specifics of the report creation process.
By solving this problem, you will get hands-on experience with the Abstract Factory design pattern, particularly in scenarios involving multiple product families (departments and report types) where creating a consistent and scalable solution is crucial.