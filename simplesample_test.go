package wkhtmltopdf

import (
	"fmt"
	"log"
)

func ExampleNewPDFGenerator() {

	// Create new PDF generator
	pdfg, err := NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Add one page from an URL
	pdfg.AddPage(NewPage("https://github.com/SebastiaanKlippert/go-wkhtmltopdf"))

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
