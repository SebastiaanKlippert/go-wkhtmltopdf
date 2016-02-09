[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf)
[![Build Status](https://travis-ci.org/SebastiaanKlippert/go-wkhtmltopdf.svg?branch=master)](https://travis-ci.org/SebastiaanKlippert/go-wkhtmltopdf)

# go-wkhtmltopdf
Golang commandline wrapper for wkhtmltopdf

Work in progress, used internally only at this point. 
No guarantees and everything may change.

See http://wkhtmltopdf.org/index.html

#What and why
go-wkhtmltopdf is a pure Golang wrapper around the wkhtmltopdf command line utility.

It has all options typed out as struct members which makes it very easy to use if you use an IDE with
code completion and it has type safety for all options.
For example you can set general options like
```go
pdfg.Dpi.Set(600)
pdfg.NoCollate.Set(false)
pdfg.PageSize.Set(PageSize_A4)
pdfg.MarginBottom.Set(40)
``` 
The same goes for adding pages, settings page options, TOC options per page etc.

It takes care of setting the correct order of options as these can become very long with muliple pages where 
you have page and TOC options for each page.

Secondly it makes usage in server-type applications easier, every instance (PDF process) has its own output buffer 
which contains the PDF output and you can feed one input document from an io.Reader (using stdin in wkhtmltopdf).
You can combine any number or external HTML documents (HTTP(S) links) with at most one HTML document from stdin and set 
options for each input document.

Note: You can also ignore the internal buffer and let wkhtmltopdf write directly to disk if required for large files.

#Installation
go get or use a Go dependency manager of your liking

```
go get -u github.com/SebastiaanKlippert/go-wkhtmltopdf
```

go-wkhtmltopdf finds the path to wkhtmltopdf by
* first looking in the current dir
* looking in the PATH and PATHEXT environment dirs
* using the WKHTMLTOPDF_PATH environment dir

The path is cached, meaning you can not change the location of wkhtmltopdf in
a running program once it has been found

#Usage
See testfile for more complex options, the most simple test is in simplesample_test.go

```go
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
	//Or to use the raw PDF data
	//pdf := pdfg.Bytes()
}

````

#Speed 
The speed if pretty much determined by wkhtmltopdf itself, or if you use extrnal source URLs, the time it takes to get the source HTML.

The go wrapper time is negligible with around 0.036ms for parsing an above average number of commandline options.

Benchmarks are included.