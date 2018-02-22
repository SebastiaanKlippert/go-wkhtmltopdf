package wkhtmltopdf

import (
	"encoding/json"
)

type jsonPDFGenerator struct {
	GlobalOptions  globalOptions
	OutlineOptions outlineOptions
	Cover          cover
	TOC            toc
}

func (pdfg *PDFGenerator) ToJSON() ([]byte, error) {

	jp := &jsonPDFGenerator{
		TOC:            pdfg.TOC,
		Cover:          pdfg.Cover,
		GlobalOptions:  pdfg.globalOptions,
		OutlineOptions: pdfg.outlineOptions,
	}

	//TODO no indents
	return json.MarshalIndent(jp, " ", " ")
}

/*
func (gopt *globalOptions) MarshalJSON() ([]byte, error) {

	return []byte{}, nil
}

func (oopt *outlineOptions) MarshalJSON() ([]byte, error) {

	return []byte{}, nil
}

func (popt *pageOptions) MarshalJSON() ([]byte, error) {

	return []byte{}, nil
}

func (topt *tocOptions) MarshalJSON() ([]byte, error) {

	return []byte{}, nil
}
*/

//TODO escape =

type jsonOption struct {
	Option string
	IsSet  bool
	Value  interface{}
}

func (bo *boolOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonOption{bo.option, true, bo.value})
}

func (so *stringOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonOption{so.option, so.value != "", so.value})
}

func (io *uintOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonOption{io.option, io.isSet, io.value})
}

func (fo *floatOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonOption{fo.option, fo.isSet, fo.value})
}

func (mo *mapOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonOption{mo.option, len(mo.value) != 0, mo.value})
}
