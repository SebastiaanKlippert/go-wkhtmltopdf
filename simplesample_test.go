package wkhtmltopdf

import (
	"fmt"
	"log"
)

func ExampleSimple() {

	pdfg, err := NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(NewPage("https://github.com/SebastiaanKlippert/go-wkhtmltopdf"))
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
	// Output: Done
}
