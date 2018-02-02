[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf)
[![Build Status](https://travis-ci.org/SebastiaanKlippert/go-wkhtmltopdf.svg?branch=master)](https://travis-ci.org/SebastiaanKlippert/go-wkhtmltopdf)
[![Go Report Card](https://goreportcard.com/badge/SebastiaanKlippert/go-wkhtmltopdf)](https://goreportcard.com/report/SebastiaanKlippert/go-wkhtmltopdf)
[![codecov](https://codecov.io/gh/SebastiaanKlippert/go-wkhtmltopdf/branch/master/graph/badge.svg)](https://codecov.io/gh/SebastiaanKlippert/go-wkhtmltopdf)

# go-wkhtmltopdf
Golang commandline wrapper for wkhtmltopdf

~~Work in progress, used internally only at this point.~~  
~~No guarantees and **everything may change**.~~  
Update 17-07-2016:  
This package is now used in our production environment after a long test period so changes are unlikely.  
Bugs wil be fixed as soon as they are found, but don't expect major changes from now on.  

See http://wkhtmltopdf.org/index.html for wkhtmltopdf docs.

# What and why
We needed a way to generate PDF documents from Go. These vary from invoices with highly customizable lay-outs to reports with tables, graphs and images. In our opinion the best way to do this was by using HTML/CSS templates as source for our PDFs. Using CSS print media types and millimeters instead of pixel units we can generate very acurate PDF documents using wkhtmltopdf.

go-wkhtmltopdf is a pure Golang wrapper around the wkhtmltopdf command line utility.

It has all options typed out as struct members which makes it very easy to use if you use an IDE with
code completion and it has type safety for all options.
For example you can set general options like
```go
pdfg.Dpi.Set(600)
pdfg.NoCollate.Set(false)
pdfg.PageSize.Set(PageSizeA4)
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

For us this is one of the easiest ways to genere PDF documents from Go(lang) and performance is very acceptable.

# Installation
go get or use a Go dependency manager of your liking.

```
go get -u github.com/SebastiaanKlippert/go-wkhtmltopdf
```

go-wkhtmltopdf finds the path to wkhtmltopdf by
* first looking in the current dir
* looking in the PATH and PATHEXT environment dirs
* using the WKHTMLTOPDF_PATH environment dir

If you need to set your own wkhtmltopdf path or want to change it during execution, you can call SetPath().

# Usage
See testfile ```wkhtmltopdf_test.go``` for more complex options, the most simple test is in ```simplesample_test.go``` 

```go
package wkhtmltopdf

import (
	"fmt"
	"log"
    "strings"
    wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)


func main(){


  pdfg, err :=  wkhtml.NewPDFGenerator()
   if err != nil{
	  return
  }
  <!-- Html String Working fine with Background Colour and Image-->
  htmlStr := `<html><body><h1 style="color:red;">This is an html from pdf to test color<h1><img src="http://api.qrserver.com/v1/create-qr-code/?data=HelloWorld" alt="img" height="42" width="42"></img></body></html>`

  pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))


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

}
```

As mentioned before, you can provide one document from stdin, this is done by using a [PageReader](https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf#PageReader "GoDoc") object as input to AddPage. This is best constructed with  [NewPageReader](https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf#NewPageReader "GoDoc") and will accept any io.Reader so this can be used with files from disk (os.File) or memory (bytes.Buffer) etc.  
A simple example snippet:
```go
	html := "<html>Hi</html>"
	pdfg, err :=  wkhtml.NewPDFGenerator()
     if err != nil{
    	  return
     }
	pdfgen.AddPage(wkhtml.NewPageReader(strings.NewReader(html)))
```

# Speed 
The speed if pretty much determined by wkhtmltopdf itself, or if you use external source URLs, the time it takes to get and render the source HTML.

The go wrapper time is negligible with around 0.04ms for parsing an above average number of commandline options.

Benchmarks are included.
