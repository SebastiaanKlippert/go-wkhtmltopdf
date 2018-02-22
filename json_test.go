package wkhtmltopdf

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPDFGenerator_ToJSON(t *testing.T) {

	pdfg := newTestPDFGenerator(t)

	jb, err := pdfg.ToJSON()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(jb))

}

func TestBoolOption_MarshalJSON(t *testing.T) {
	bo := &boolOption{"option", true}
	assertJSON(t, bo, "")
}

func assertJSON(t *testing.T, option interface{}, want string) {
	j, err := json.Marshal(option)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(j))
}
