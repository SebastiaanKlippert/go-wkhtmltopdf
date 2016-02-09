//Package wkhtmltopdf contains wrappers around the wkhtmltopdf commandline tool
package wkhtmltopdf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//for each page
type Page struct {
	Input string
	PageOptions
}

func (p *Page) InputFile() string {
	return p.Input
}

func (p *Page) Args() []string {
	return p.PageOptions.Args()
}

func (pr *Page) Reader() io.Reader {
	return nil
}

func NewPage(input string) *Page {
	return &Page{
		Input:       input,
		PageOptions: NewPageOptions(),
	}
}

//PageReader is one input page (a HTML document) that is read from an io.Reader
//You can add only one Page from a reader
type PageReader struct {
	Input io.Reader
	PageOptions
}

func (pr *PageReader) InputFile() string {
	return "-"
}

func (pr *PageReader) Args() []string {
	return pr.PageOptions.Args()
}

func (pr *PageReader) Reader() io.Reader {
	return pr.Input
}

func NewPageReader(input io.Reader) *PageReader {
	return &PageReader{
		Input:       input,
		PageOptions: NewPageOptions(),
	}
}

type page interface {
	Args() []string
	InputFile() string
	Reader() io.Reader
}

type PageOptions struct {
	pageOptions
	headerAndFooterOptions
}

func (po *PageOptions) Args() []string {
	return append(append([]string{}, po.pageOptions.Args()...), po.headerAndFooterOptions.Args()...)
}

func NewPageOptions() PageOptions {
	return PageOptions{
		pageOptions:            newPageOptions(),
		headerAndFooterOptions: newHeaderAndFooterOptions(),
	}
}

//cover page
type cover struct {
	Input string
	pageOptions
}

//table of contents
type toc struct {
	Include bool
	allTocOptions
}

type allTocOptions struct {
	pageOptions
	tocOptions
}

//PdfGenerator is the main wkhtmltopdf struct
type PdfGenerator struct {
	globalOptions
	outlineOptions

	Cover      cover
	TOC        toc
	OutputFile string //filename to write to, default empty (writes to internal buffer)

	binPath string
	outbuf  bytes.Buffer
	pages   []page
}

//Args returns the commandline arguments as a string slice
func (pdfg *PdfGenerator) Args() []string {
	args := []string{}
	args = append(args, pdfg.globalOptions.Args()...)
	args = append(args, pdfg.outlineOptions.Args()...)
	if pdfg.Cover.Input != "" {
		args = append(args, "cover")
		args = append(args, pdfg.Cover.Input)
		args = append(args, pdfg.Cover.pageOptions.Args()...)
	}
	if pdfg.TOC.Include {
		args = append(args, "toc")
		args = append(args, pdfg.TOC.pageOptions.Args()...)
		args = append(args, pdfg.TOC.tocOptions.Args()...)
	}
	for _, page := range pdfg.pages {
		args = append(args, "page")
		args = append(args, page.InputFile())
		args = append(args, page.Args()...)
	}
	if pdfg.OutputFile != "" {
		args = append(args, pdfg.OutputFile)
	} else {
		args = append(args, "-")
	}
	return args
}

//Argstring returns Args as a single string
func (pdfg *PdfGenerator) ArgString() string {
	return strings.Join(pdfg.Args(), " ")
}

//AddPage adds a new input page to the document.
//A page is an input HTML page, it can span multiple pages in the output document.
//It is a Page when read from file or URL or a PageReader when read from memory.
func (pdfg *PdfGenerator) AddPage(p page) {
	pdfg.pages = append(pdfg.pages, p)
}

//SetPages resets all pages
func (pdfg *PdfGenerator) SetPages(p []page) {
	pdfg.pages = p
}

//Buffer returns the embedded output buffer used if OutputFile is empty
func (pdfg *PdfGenerator) Buffer() *bytes.Buffer {
	return &pdfg.outbuf
}

//Bytes returns the output byte slice from the output buffer used if OutputFile is empty
func (pdfg *PdfGenerator) Bytes() []byte {
	return pdfg.outbuf.Bytes()
}

//WriteFile writes the contents of the output buffer to a file
func (pdfg *PdfGenerator) WriteFile(filename string) error {
	return ioutil.WriteFile(filename, pdfg.Bytes(), os.ModeExclusive)
}

//findPath finds the path to wkhtmltopdf by
//- first looking in the current dir
//- looking in the PATH and PATHEXT environment dirs
//- using the WKHTMLTOPDF_PATH environment dir
//The path is cached, meaning you can not change the location of wkhtmltopdf in
//a running program once it has been found
func (pdfg *PdfGenerator) findPath() error {
	const exe = "wkhtmltopdf"
	if binPath != "" {
		pdfg.binPath = binPath
		return nil
	}
	exe_dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	path, err := exec.LookPath(filepath.Join(exe_dir, exe))
	if err == nil && path != "" {
		binPath = path
		pdfg.binPath = path
		return nil
	}
	path, err = exec.LookPath(exe)
	if err == nil && path != "" {
		binPath = path
		pdfg.binPath = path
		return nil
	}
	dir := os.Getenv("WKHTMLTOPDF_PATH")
	if dir == "" {
		return fmt.Errorf("%s not found", exe)
	}
	path, err = exec.LookPath(filepath.Join(dir, exe))
	if err == nil && path != "" {
		binPath = path
		pdfg.binPath = path
		return nil
	}
	return fmt.Errorf("%s not found", exe)
}

var binPath string //the cached paths as used by findPath()

func (pdfg *PdfGenerator) Create() error {
	return pdfg.run()
}

func (pdfg *PdfGenerator) run() error {

	errbuf := &bytes.Buffer{}

	cmd := exec.Command(pdfg.binPath, pdfg.Args()...)

	cmd.Stdout = &pdfg.outbuf
	cmd.Stderr = errbuf
	//if there is a pageReader page (from Stdin) we set Stdin to that reader
	for _, page := range pdfg.pages {
		if page.Reader() != nil {
			cmd.Stdin = page.Reader()
			break
		}
	}

	err := cmd.Run()
	if err != nil {
		//return first line only
		bs := bufio.NewScanner(errbuf)
		if bs.Scan() {
			return errors.New(bs.Text())
		}
		return err
	}
	return nil
}

//NewPDFGenerator returns a new wkhtmltopdf struct
func NewPDFGenerator() (*PdfGenerator, error) {
	pdfg := &PdfGenerator{
		globalOptions:  newGlobalOptions(),
		outlineOptions: newOutlineOptions(),
		Cover: cover{
			pageOptions: newPageOptions(),
		},
		TOC: toc{
			allTocOptions: allTocOptions{
				tocOptions:  newTocOptions(),
				pageOptions: newPageOptions(),
			},
		},
	}
	err := pdfg.findPath()
	return pdfg, err
}
