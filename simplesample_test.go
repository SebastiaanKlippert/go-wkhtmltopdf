package wkhtmltopdf

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

func ExampleNewPDFGenerator() {

	// Create new PDF generator
	pdfg, err := NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Add one page from an URL
	pdfg.AddPage(NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf"))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done

}

func ExampleNewPDFGeneratorFromJSON() {

	const html = `<!doctype html><html><head><title>WKHTMLTOPDF TEST</title></head><body>HELLO PDF</body></html>`

	// Client code
	pdfg := NewPDFPreparer()
	pdfg.AddPage(NewPageReader(strings.NewReader(html)))
	pdfg.Dpi.Set(600)

	// The html string is also saved as base64 string in the JSON file
	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	// The JSON can be saved, uploaded, etc.

	// Server code, create a new PDF generator from JSON, also looks for the wkhtmltopdf executable
	pdfgFromJSON, err := NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Create the PDF
	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Use the PDF
	fmt.Printf("PDF size %d bytes", pdfgFromJSON.Buffer().Len())
}
