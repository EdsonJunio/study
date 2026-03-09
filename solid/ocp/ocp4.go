package ocp

import "fmt"

type (
	Report struct {
		Data []string
	}

	Exporter interface {
		Export(r *Report) error
	}

	ReportCSV struct{}

	ReportPDF struct{}

	ReportJSON struct{}

	ReportService struct{}
)

func (r *ReportCSV) Export(re *Report) error {
	fmt.Println("Exporting data in CSV format (comma-separated)...")
	for _, linha := range re.Data {
		fmt.Println(linha)
	}

	return nil
}

func (r *ReportPDF) Export(re *Report) error {
	fmt.Println("Exporting data in PDF format (with tables and logo)...")
	for _, linha := range re.Data {
		fmt.Println(linha)
	}

	return nil
}

func (r *ReportJSON) Export(re *Report) error {
	fmt.Println("Exporting data in JSON format...")

	return nil
}

func (r *ReportService) Export(re *Report, e Exporter) error {
	err := e.Export(re)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	export := &Report{
		Data: []string{"This is an export to test", "hello"},
	}

	service := ReportService{}
	err := service.Export(export, &ReportCSV{})
	if err != nil {
		fmt.Println("Error:", err)
	}

}
