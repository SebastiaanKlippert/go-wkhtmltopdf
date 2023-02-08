package wkhtmltopdf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPDFGenerator_ToJSON(t *testing.T) {
	pdfg := newTestPDFGenerator(t)

	// add a reader page as well
	htmlfile, err := os.Open("testdata/htmlsimple.html")
	if err != nil {
		t.Fatal(err)
	}
	defer htmlfile.Close()

	pdfg.AddPage(NewPageReader(htmlfile))

	jb, err := pdfg.ToJSON()
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile("testdata/expected.json")
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(t, string(expected), string(jb))
}

func TestNewPDFGeneratorFromJSON(t *testing.T) {
	pdfg := newTestPDFGenerator(t)
	jb, err := pdfg.ToJSON()
	if err != nil {
		t.Fatal(err)
	}

	pdfgFromJSON, err := NewPDFGeneratorFromJSON(bytes.NewReader(jb))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedArgString(), pdfgFromJSON.ArgString())

	pdfgFromJSON.OutputFile = "testdata/TestNewPDFGeneratorFromJSON.PDF"

	err = pdfgFromJSON.Create()
	if err != nil {
		t.Fatal(err)
	}

}

func TestNewPDFGeneratorFromJSONWithReader(t *testing.T) {

	pdfg := NewPDFPreparer()
	htmlfile, err := ioutil.ReadFile("testdata/htmlsimple.html")
	if err != nil {
		t.Fatal(err)
	}
	pdfg.AddPage(NewPageReader(bytes.NewReader(htmlfile)))

	jb, err := pdfg.ToJSON()
	if err != nil {
		t.Fatal(err)
	}

	pdfgFromJSON, err := NewPDFGeneratorFromJSON(bytes.NewReader(jb))
	if err != nil {
		t.Fatal(err)
	}

	// assert argstring
	if pdfg.ArgString() != pdfgFromJSON.ArgString() {
		t.Errorf("Want argstring:\n%s\nHave:\n%s", pdfg.ArgString(), pdfgFromJSON.ArgString())
	}

	// assert content
	buf, err := ioutil.ReadAll(pdfgFromJSON.pages[0].Reader())
	if err != nil {
		t.Fatal(err)
	}
	if string(buf) != string(htmlfile) {
		t.Errorf("Want HTML:\n%s\nHave:\n%s", string(htmlfile), string(buf))
	}

}

func TestBoolOption_JSON(t *testing.T) {
	bo := &boolOption{"option", true}
	assertJSON(t, bo, new(boolOption))
}

func TestFloatOptionJSON(t *testing.T) {
	fo := &floatOption{"option", 1.11, true}
	assertJSON(t, fo, new(floatOption))
}

func TestMapOption_JSON(t *testing.T) {
	mo := &mapOption{"option", map[string]string{"foo1": "bar1", "foo2": "bar2"}}
	assertJSON(t, mo, new(mapOption))
}

func TestUintOption_JSON(t *testing.T) {
	uo := &uintOption{"option", 111, true}
	assertJSON(t, uo, new(uintOption))
}

func TestStringOption_JSON(t *testing.T) {
	so := &stringOption{"option", "abc"}
	assertJSON(t, so, new(stringOption))
}

func TestSliceOption_JSON(t *testing.T) {
	so := &sliceOption{"option", []string{"foo", "bar"}}
	assertJSON(t, so, new(sliceOption))
}

func assertJSON(t *testing.T, option, newOption interface{}) {
	j, err := json.Marshal(option)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(j, newOption)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(option, newOption) {
		t.Errorf("Diff after marshal and unmarshal:\n%+v\n%+v", option, newOption)
	}
}
