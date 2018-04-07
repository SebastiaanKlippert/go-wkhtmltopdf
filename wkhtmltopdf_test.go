package wkhtmltopdf

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"strings"
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
	pdfg.MarginLeft.Set(0)

	page1 := NewPage("https://www.google.com")

	page1.DisableSmartShrinking.Set(true)
	page1.HeaderSpacing.Set(10.01)
	page1.Allow.Set("/usr/local/html")
	page1.Allow.Set("/usr/local/images")
	page1.CustomHeader.Set("X-AppKey", "abcdef")
	page1.ViewportSize.Set("3840x2160")

	if runtime.GOOS == "darwin" {
		page1.LoadErrorHandling.Set("ignore")
	}

	pdfg.AddPage(page1)

	pdfg.Cover.Input = "https://wkhtmltopdf.org/index.html"
	pdfg.Cover.Zoom.Set(0.75)

	pdfg.TOC.Include = true
	pdfg.TOC.DisableDottedLines.Set(true)

	return pdfg
}

func wantArgString() string {
	return "--dpi 600 --margin-bottom 40 --margin-left 0 --page-size A4 cover https://wkhtmltopdf.org/index.html --zoom 0.750 toc --disable-dotted-lines page https://www.google.com --allow /usr/local/html --allow /usr/local/images --custom-header X-AppKey abcdef --disable-smart-shrinking --viewport-size 3840x2160 --header-spacing 10.010 -"
}

func TestArgString(t *testing.T) {
	pdfg := newTestPDFGenerator(t)
	want := wantArgString()
	if pdfg.ArgString() != want {
		t.Errorf("Want argstring:\n%s\nHave:\n%s", want, pdfg.ArgString())
	}
	pdfg.SetPages(pdfg.pages)
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
	//TODO temp error check because older versions of wkhtmltopdf return a different error :(
	wantErrNew := "You need to specify at least one input file, and exactly one output file"
	wantErrOld := "You need to specify atleast one input file, and exactly one output file"
	if strings.HasPrefix(err.Error(), wantErrNew) == false && strings.HasPrefix(err.Error(), wantErrOld) == false {
		t.Errorf("Want error prefix %s or %s, have %s", wantErrNew, wantErrOld, err.Error())
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

func TestGeneratePdfFromStdinSimple(t *testing.T) {
	//Use a new blank PDF generator
	pdfg, err := NewPDFGenerator()
	if err != nil {
		t.Fatal(err)
	}
	htmlfile, err := ioutil.ReadFile("./testfiles/htmlsimple.html")
	if err != nil {
		t.Fatal(err)
	}
	pdfg.AddPage(NewPageReader(bytes.NewReader(htmlfile)))
	err = pdfg.Create()
	if err != nil {
		t.Fatal(err)
	}
	err = pdfg.WriteFile("./testfiles/TestGeneratePdfFromStdinSimple.pdf")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PDF size %vkB", len(pdfg.Bytes())/1024)
	if pdfg.Buffer().Len() != len(pdfg.Bytes()) {
		t.Errorf("Buffersize not equal")
	}
}

func TestPDFGeneratorOutputFile(t *testing.T) {
	pdfg, err := NewPDFGenerator()
	if err != nil {
		t.Fatal(err)
	}
	htmlfile, err := os.Open("./testfiles/htmlsimple.html")
	if err != nil {
		t.Fatal(err)
	}
	defer htmlfile.Close()

	pdfg.OutputFile = "./testfiles/TestPDFGeneratorOutputFile.pdf"

	pdfg.AddPage(NewPageReader(htmlfile))
	err = pdfg.Create()
	if err != nil {
		t.Fatal(err)
	}

	pdfFile, err := os.Open("./testfiles/TestPDFGeneratorOutputFile.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer pdfFile.Close()

	stat, err := pdfFile.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if stat.Size() < 100 {
		t.Errorf("generated PDF is size under 100 bytes")
	}
}

func TestGeneratePdfFromStdinHtml5(t *testing.T) {
	//Use newTestPDFGenerator and append to page1 and TOC
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
	err = pdfg.WriteFile("./testfiles/TestGeneratePdfFromStdinHtml5.pdf")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PDF size %vkB", len(pdfg.Bytes())/1024)
}

func TestPath(t *testing.T) {
	path := "/usr/wkhtmltopdf/wkhtmltopdf"
	SetPath(path)
	defer SetPath("")
	if GetPath() != path {
		t.Errorf("Have path %q, want %q", GetPath(), path)
	}
}

func BenchmarkArgs(b *testing.B) {
	pdfg := newTestPDFGenerator(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pdfg.Args()
	}
}

func TestStringOption(t *testing.T) {
	opt := stringOption{
		option: "stringopt",
	}
	opt.Set("value123")

	want := []string{"--stringopt", "value123"}

	if !reflect.DeepEqual(opt.Parse(), want) {
		t.Errorf("expected %v, have %v", want, opt.Parse())
	}

}

func TestSliceOption(t *testing.T) {
	opt := sliceOption{
		option: "sliceopt",
	}
	opt.Set("string15183")
	opt.Set("foo")
	opt.Set("bar")

	want := []string{"--sliceopt", "string15183", "--sliceopt", "foo", "--sliceopt", "bar"}

	if !reflect.DeepEqual(opt.Parse(), want) {
		t.Errorf("expected %v, have %v", want, opt.Parse())
	}

}

func TestMapOption(t *testing.T) {
	opt := mapOption{
		option: "mapopt",
	}

	opt.Set("key1", "foo")
	opt.Set("key2", "bar")
	opt.Set("key3", "Hello")

	result := strings.Join(opt.Parse(), " ")
	if !strings.Contains(result, "--mapopt key1 foo") {
		t.Error("missing map option key1")
	}
	if !strings.Contains(result, "--mapopt key2 bar") {
		t.Error("missing map option key2")
	}
	if !strings.Contains(result, "--mapopt key3 Hello") {
		t.Error("missing map option key3")
	}
}

func TestUIntOption(t *testing.T) {
	opt := uintOption{
		option: "uintopt",
	}
	opt.Set(14860)

	want := []string{"--uintopt", "14860"}

	if !reflect.DeepEqual(opt.Parse(), want) {
		t.Errorf("expected %v, have %v", want, opt.Parse())
	}
}

func TestFloatOption(t *testing.T) {
	opt := floatOption{
		option: "flopt",
	}
	opt.Set(239.75)

	want := []string{"--flopt", "239.750"}

	if !reflect.DeepEqual(opt.Parse(), want) {
		t.Errorf("expected %v, have %v", want, opt.Parse())
	}
}

func TestBoolOption(t *testing.T) {
	opt := boolOption{
		option: "boolopt",
	}
	opt.Set(true)

	want := []string{"--boolopt"}

	if !reflect.DeepEqual(opt.Parse(), want) {
		t.Errorf("expected %v, have %v", want, opt.Parse())
	}
}
