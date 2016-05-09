package wkhtmltopdf

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func newTestPDFGenerator(tb testing.TB) *PDFGenerator {

	pdfg, err := NewPDFGenerator()
	if err != nil {
		tb.Fatal(err)
	}

	pdfg.Dpi.Set(600)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(PageSizeA4)
	pdfg.MarginBottom.Set(40)

	page1 := NewPage("https://www.google.com")

	page1.DisableSmartShrinking.Set(true)
	page1.HeaderSpacing.Set(10.01)
	page1.Allow.Set("/usr/local/html")
	page1.Allow.Set("/usr/local/images")
	page1.CustomHeader.Set("X-AppKey", "abcdef")

	pdfg.AddPage(page1)

	pdfg.Cover.Input = "http://wkhtmltopdf.org/index.html"
	pdfg.Cover.Zoom.Set(0.75)

	pdfg.TOC.Include = true
	pdfg.TOC.DisableDottedLines.Set(true)

	return pdfg
}

func TestArgString(t *testing.T) {
	pdfg := newTestPDFGenerator(t)
	want := "--dpi 600 --margin-bottom 40 --page-size A4 cover http://wkhtmltopdf.org/index.html --zoom 0.750 toc --disable-dotted-lines page https://www.google.com --allow /usr/local/html --allow /usr/local/images --custom-header X-AppKey abcdef --disable-smart-shrinking --header-spacing 10.010 -"
	if pdfg.ArgString() != want {
		t.Errorf("Want argstring:\n%s\nHave:\n%s", want, pdfg.ArgString())
	}
}

func TestVersion(t *testing.T) {
	pdfg, err := NewPDFGenerator()
	if err != nil {
		t.Fatal(err)
	}
	pdfg.Version.Set(true)
	err = pdfg.Create()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoInput(t *testing.T) {
	pdfg, err := NewPDFGenerator()
	if err != nil {
		t.Fatal(err)
	}
	err = pdfg.Create()
	if err == nil {
		t.Fatal("Want an error when there is no input, have no error")
	}
	wantErr := "You need to specify atleast one input file, and exactly one output file"
	if err.Error() != wantErr {
		t.Errorf("Want error %s, have %s", wantErr, err.Error())
	}
}

func TestGeneratePDF(t *testing.T) {
	pdfg := newTestPDFGenerator(t)
	err := pdfg.Create()
	if err != nil {
		t.Fatal(err)
	}
	err = pdfg.WriteFile("./testfiles/TestGeneratePDF.pdf")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PDF size %vkB", len(pdfg.Bytes())/1024)
}

func TestGeneratePdfFromStdin(t *testing.T) {
	pdfg := newTestPDFGenerator(t)
	htmlfile, err := ioutil.ReadFile("./testfiles/html5.html")
	if err != nil {
		t.Fatal(err)
	}
	page2 := NewPageReader(bytes.NewReader(htmlfile))
	pdfg.AddPage(page2)
	err = pdfg.Create()
	if err != nil {
		t.Fatal(err)
	}
	err = pdfg.WriteFile("./testfiles/TestGeneratePdfFromStdin.pdf")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PDF size %vkB", len(pdfg.Bytes())/1024)
}

func BenchmarkArgs(b *testing.B) {
	pdfg := newTestPDFGenerator(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pdfg.Args()
	}
}
