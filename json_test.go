package wkhtmltopdf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestPDFGenerator_ToJSON(t *testing.T) {
	pdfg := newTestPDFGenerator(t)

	// add a reader page as well
	htmlfile, err := os.Open("./testfiles/htmlsimple.html")
	if err != nil {
		t.Fatal(err)
	}
	defer htmlfile.Close()

	pdfg.AddPage(NewPageReader(htmlfile))

	jb, err := pdfg.ToJSON()
	if err != nil {
		t.Fatal(err)
	}

	l := 15230
	if len(jb) != l {
		t.Errorf("Want %d JSON bytes, have %d", l, len(jb))
	}

	want := `{"GlobalOptions":{"CookieJar":{"Option":"cookie-jar","Value":""},"Copies":{"Option":"copies","IsSet":false,"Value":0},"Dpi":{"Option":"dpi","IsSet":true,"Value":600},"ExtendedHelp":{"Option":"extended-help","Value":false},"Grayscale":{"Option":"grayscale","Value":false},"Help":{"Option":"true","Value":false},"HTMLDoc":{"Option":"htmldoc","Value":false},"ImageDpi":{"Option":"image-dpi","IsSet":false,"Value":0},"ImageQuality":{"Option":"image-quality","IsSet":false,"Value":0},"License":{"Option":"license","Value":false},"Lowquality":{"Option":"lowquality","Value":false},"ManPage":{"Option":"manpage","Value":false},"MarginBottom":{"Option":"margin-bottom","IsSet":true,"Value":40},"MarginLeft":{"Option":"margin-left","IsSet":true,"Value":0},"MarginRight":{"Option":"margin-right","IsSet":false,"Value":0},"MarginTop":{"Option":"margin-top","IsSet":false,"Value":0},"NoCollate":{"Option":"nocollate","Value":false},"NoPdfCompression":{"Option":"no-pdf-compression","Value":false},"Orientation":{"Option":"orientation","Value":""},"PageHeight":{"Option":"page-height","IsSet":false,"Value":0},"PageSize":{"Option":"page-size","Value":"A4"},"PageWidth":{"Option":"page-width","IsSet":false,"Value":0},"Quiet":{"Option":"quiet","Value":false},"ReadArgsFromStdin":{"Option":"read-args-from-stdin","Value":false},"Readme":{"Option":"readme","Value":false},"Title":{"Option":"title","Value":""},"Version":{"Option":"version","Value":false}},"OutlineOptions":{"DumpDefaultTocXsl":{"Option":"dump-default-toc-xsl","Value":false},"DumpOutline":{"Option":"dump-outline","Value":""},"NoOutline":{"Option":"no-outline","Value":false},"OutlineDepth":{"Option":"outline-depth","IsSet":false,"Value":0}},"Cover":{"Input":"https://wkhtmltopdf.org/index.html","Allow":{"Option":"allow","Value":null},"BypassProxyFor":{"Option":"bypass-proxy-for","Value":null},"CacheDir":{"Option":"cache-dir","Value":""},"CheckboxCheckedSvg":{"Option":"checkbox-checked-svg","Value":""},"CheckboxSvg":{"Option":"checkbox-svg","Value":""},"Cookie":{"Option":"cookie","Value":null},"CustomHeader":{"Option":"custom-header","Value":null},"CustomHeaderPropagation":{"Option":"custom-header-propagation","Value":false},"DebugJavascript":{"Option":"debug-javascript","Value":false},"DefaultHeader":{"Option":"default-header","Value":false},"DisableExternalLinks":{"Option":"disable-external-links","Value":false},"DisableInternalLinks":{"Option":"disable-internal-links","Value":false},"DisableJavascript":{"Option":"disable-javascript","Value":false},"DisableLocalFileAccess":{"Option":"disable-local-file-access","Value":false},"DisableSmartShrinking":{"Option":"disable-smart-shrinking","Value":false},"EnableForms":{"Option":"enable-forms","Value":false},"EnablePlugins":{"Option":"enable-plugins","Value":false},"EnableTocBackLinks":{"Option":"enable-toc-back-links","Value":false},"Encoding":{"Option":"encoding","Value":""},"ExcludeFromOutline":{"Option":"exclude-from-outline","Value":false},"JavascriptDelay":{"Option":"javascript-delay","IsSet":false,"Value":0},"KeepRelativeLinks":{"Option":"keep-relative-links","Value":false},"LoadErrorHandling":{"Option":"load-error-handling","Value":""},"LoadMediaErrorHandling":{"Option":"load-media-error-handling","Value":""},"MinimumFontSize":{"Option":"minimum-font-size","IsSet":false,"Value":0},"NoBackground":{"Option":"no-background","Value":false},"NoCustomHeaderPropagation":{"Option":"no-custom-header-propagation","Value":false},"NoImages":{"Option":"no-images","Value":false},"NoStopSlowScripts":{"Option":"no-stop-slow-scripts","Value":false},"PageOffset":{"Option":"page-offset","IsSet":false,"Value":0},"Password":{"Option":"password","Value":""},"Post":{"Option":"post","Value":null},"PostFile":{"Option":"post-file","Value":null},"PrintMediaType":{"Option":"print-media-type","Value":false},"Proxy":{"Option":"proxy","Value":""},"RadiobuttonCheckedSvg":{"Option":"radiobutton-checked-svg","Value":""},"RadiobuttonSvg":{"Option":"radiobutton-svg","Value":""},"RunScript":{"Option":"run-script","Value":null},"SslCrtPath":{"Option":"ssl-crt-path","Value":""},"SslKeyPassword":{"Option":"ssl-key-password","Value":""},"SslKeyPath":{"Option":"ssl-key-path","Value":""},"Username":{"Option":"username","Value":""},"UserStyleSheet":{"Option":"user-style-sheet","Value":""},"ViewportSize":{"Option":"viewport-size","Value":""},"WindowStatus":{"Option":"window-status","Value":""},"Zoom":{"Option":"zoom","IsSet":true,"Value":0.75}},"TOC":{"Include":true,"Allow":{"Option":"allow","Value":null},"BypassProxyFor":{"Option":"bypass-proxy-for","Value":null},"CacheDir":{"Option":"cache-dir","Value":""},"CheckboxCheckedSvg":{"Option":"checkbox-checked-svg","Value":""},"CheckboxSvg":{"Option":"checkbox-svg","Value":""},"Cookie":{"Option":"cookie","Value":null},"CustomHeader":{"Option":"custom-header","Value":null},"CustomHeaderPropagation":{"Option":"custom-header-propagation","Value":false},"DebugJavascript":{"Option":"debug-javascript","Value":false},"DefaultHeader":{"Option":"default-header","Value":false},"DisableExternalLinks":{"Option":"disable-external-links","Value":false},"DisableInternalLinks":{"Option":"disable-internal-links","Value":false},"DisableJavascript":{"Option":"disable-javascript","Value":false},"DisableLocalFileAccess":{"Option":"disable-local-file-access","Value":false},"DisableSmartShrinking":{"Option":"disable-smart-shrinking","Value":false},"EnableForms":{"Option":"enable-forms","Value":false},"EnablePlugins":{"Option":"enable-plugins","Value":false},"EnableTocBackLinks":{"Option":"enable-toc-back-links","Value":false},"Encoding":{"Option":"encoding","Value":""},"ExcludeFromOutline":{"Option":"exclude-from-outline","Value":false},"JavascriptDelay":{"Option":"javascript-delay","IsSet":false,"Value":0},"KeepRelativeLinks":{"Option":"keep-relative-links","Value":false},"LoadErrorHandling":{"Option":"load-error-handling","Value":""},"LoadMediaErrorHandling":{"Option":"load-media-error-handling","Value":""},"MinimumFontSize":{"Option":"minimum-font-size","IsSet":false,"Value":0},"NoBackground":{"Option":"no-background","Value":false},"NoCustomHeaderPropagation":{"Option":"no-custom-header-propagation","Value":false},"NoImages":{"Option":"no-images","Value":false},"NoStopSlowScripts":{"Option":"no-stop-slow-scripts","Value":false},"PageOffset":{"Option":"page-offset","IsSet":false,"Value":0},"Password":{"Option":"password","Value":""},"Post":{"Option":"post","Value":null},"PostFile":{"Option":"post-file","Value":null},"PrintMediaType":{"Option":"print-media-type","Value":false},"Proxy":{"Option":"proxy","Value":""},"RadiobuttonCheckedSvg":{"Option":"radiobutton-checked-svg","Value":""},"RadiobuttonSvg":{"Option":"radiobutton-svg","Value":""},"RunScript":{"Option":"run-script","Value":null},"SslCrtPath":{"Option":"ssl-crt-path","Value":""},"SslKeyPassword":{"Option":"ssl-key-password","Value":""},"SslKeyPath":{"Option":"ssl-key-path","Value":""},"Username":{"Option":"username","Value":""},"UserStyleSheet":{"Option":"user-style-sheet","Value":""},"ViewportSize":{"Option":"viewport-size","Value":""},"WindowStatus":{"Option":"window-status","Value":""},"Zoom":{"Option":"zoom","IsSet":false,"Value":0},"DisableDottedLines":{"Option":"disable-dotted-lines","Value":true},"DisableTocLinks":{"Option":"disable-toc-links","Value":false},"TocHeaderText":{"Option":"toc-header-text","Value":""},"TocLevelIndentation":{"Option":"toc-level-indentation","IsSet":false,"Value":0},"TocTextSizeShrink":{"Option":"toc-text-size-shrink","IsSet":false,"Value":0},"XslStyleSheet":{"Option":"xsl-style-sheet","Value":""}},"Pages":[{"PageOptions":{"Allow":{"Option":"allow","Value":["/usr/local/html","/usr/local/images"]},"BypassProxyFor":{"Option":"bypass-proxy-for","Value":null},"CacheDir":{"Option":"cache-dir","Value":""},"CheckboxCheckedSvg":{"Option":"checkbox-checked-svg","Value":""},"CheckboxSvg":{"Option":"checkbox-svg","Value":""},"Cookie":{"Option":"cookie","Value":null},"CustomHeader":{"Option":"custom-header","Value":{"X-AppKey":"abcdef"}},"CustomHeaderPropagation":{"Option":"custom-header-propagation","Value":false},"DebugJavascript":{"Option":"debug-javascript","Value":false},"DefaultHeader":{"Option":"default-header","Value":false},"DisableExternalLinks":{"Option":"disable-external-links","Value":false},"DisableInternalLinks":{"Option":"disable-internal-links","Value":false},"DisableJavascript":{"Option":"disable-javascript","Value":false},"DisableLocalFileAccess":{"Option":"disable-local-file-access","Value":false},"DisableSmartShrinking":{"Option":"disable-smart-shrinking","Value":true},"EnableForms":{"Option":"enable-forms","Value":false},"EnablePlugins":{"Option":"enable-plugins","Value":false},"EnableTocBackLinks":{"Option":"enable-toc-back-links","Value":false},"Encoding":{"Option":"encoding","Value":""},"ExcludeFromOutline":{"Option":"exclude-from-outline","Value":false},"JavascriptDelay":{"Option":"javascript-delay","IsSet":false,"Value":0},"KeepRelativeLinks":{"Option":"keep-relative-links","Value":false},"LoadErrorHandling":{"Option":"load-error-handling","Value":""},"LoadMediaErrorHandling":{"Option":"load-media-error-handling","Value":""},"MinimumFontSize":{"Option":"minimum-font-size","IsSet":false,"Value":0},"NoBackground":{"Option":"no-background","Value":false},"NoCustomHeaderPropagation":{"Option":"no-custom-header-propagation","Value":false},"NoImages":{"Option":"no-images","Value":false},"NoStopSlowScripts":{"Option":"no-stop-slow-scripts","Value":false},"PageOffset":{"Option":"page-offset","IsSet":false,"Value":0},"Password":{"Option":"password","Value":""},"Post":{"Option":"post","Value":null},"PostFile":{"Option":"post-file","Value":null},"PrintMediaType":{"Option":"print-media-type","Value":false},"Proxy":{"Option":"proxy","Value":""},"RadiobuttonCheckedSvg":{"Option":"radiobutton-checked-svg","Value":""},"RadiobuttonSvg":{"Option":"radiobutton-svg","Value":""},"RunScript":{"Option":"run-script","Value":null},"SslCrtPath":{"Option":"ssl-crt-path","Value":""},"SslKeyPassword":{"Option":"ssl-key-password","Value":""},"SslKeyPath":{"Option":"ssl-key-path","Value":""},"Username":{"Option":"username","Value":""},"UserStyleSheet":{"Option":"user-style-sheet","Value":""},"ViewportSize":{"Option":"viewport-size","Value":"3840x2160"},"WindowStatus":{"Option":"window-status","Value":""},"Zoom":{"Option":"zoom","IsSet":false,"Value":0},"FooterCenter":{"Option":"footer-center","Value":""},"FooterFontName":{"Option":"footer-font-name","Value":""},"FooterFontSize":{"Option":"footer-font-size","IsSet":false,"Value":0},"FooterHTML":{"Option":"footer-html","Value":""},"FooterLeft":{"Option":"footer-left","Value":""},"FooterLine":{"Option":"footer-line","Value":false},"FooterRight":{"Option":"footer-right","Value":""},"FooterSpacing":{"Option":"footer-spacing","IsSet":false,"Value":0},"HeaderCenter":{"Option":"header-center","Value":""},"HeaderFontName":{"Option":"header-font-name","Value":""},"HeaderFontSize":{"Option":"header-font-size","IsSet":false,"Value":0},"HeaderHTML":{"Option":"header-html","Value":""},"HeaderLeft":{"Option":"header-left","Value":""},"HeaderLine":{"Option":"header-line","Value":false},"HeaderRight":{"Option":"header-right","Value":""},"HeaderSpacing":{"Option":"header-spacing","IsSet":true,"Value":10.01},"Replace":{"Option":"replace","Value":null}},"InputFile":"https://www.google.com","Base64PageData":""},{"PageOptions":{"Allow":{"Option":"allow","Value":null},"BypassProxyFor":{"Option":"bypass-proxy-for","Value":null},"CacheDir":{"Option":"cache-dir","Value":""},"CheckboxCheckedSvg":{"Option":"checkbox-checked-svg","Value":""},"CheckboxSvg":{"Option":"checkbox-svg","Value":""},"Cookie":{"Option":"cookie","Value":null},"CustomHeader":{"Option":"custom-header","Value":null},"CustomHeaderPropagation":{"Option":"custom-header-propagation","Value":false},"DebugJavascript":{"Option":"debug-javascript","Value":false},"DefaultHeader":{"Option":"default-header","Value":false},"DisableExternalLinks":{"Option":"disable-external-links","Value":false},"DisableInternalLinks":{"Option":"disable-internal-links","Value":false},"DisableJavascript":{"Option":"disable-javascript","Value":false},"DisableLocalFileAccess":{"Option":"disable-local-file-access","Value":false},"DisableSmartShrinking":{"Option":"disable-smart-shrinking","Value":false},"EnableForms":{"Option":"enable-forms","Value":false},"EnablePlugins":{"Option":"enable-plugins","Value":false},"EnableTocBackLinks":{"Option":"enable-toc-back-links","Value":false},"Encoding":{"Option":"encoding","Value":""},"ExcludeFromOutline":{"Option":"exclude-from-outline","Value":false},"JavascriptDelay":{"Option":"javascript-delay","IsSet":false,"Value":0},"KeepRelativeLinks":{"Option":"keep-relative-links","Value":false},"LoadErrorHandling":{"Option":"load-error-handling","Value":""},"LoadMediaErrorHandling":{"Option":"load-media-error-handling","Value":""},"MinimumFontSize":{"Option":"minimum-font-size","IsSet":false,"Value":0},"NoBackground":{"Option":"no-background","Value":false},"NoCustomHeaderPropagation":{"Option":"no-custom-header-propagation","Value":false},"NoImages":{"Option":"no-images","Value":false},"NoStopSlowScripts":{"Option":"no-stop-slow-scripts","Value":false},"PageOffset":{"Option":"page-offset","IsSet":false,"Value":0},"Password":{"Option":"password","Value":""},"Post":{"Option":"post","Value":null},"PostFile":{"Option":"post-file","Value":null},"PrintMediaType":{"Option":"print-media-type","Value":false},"Proxy":{"Option":"proxy","Value":""},"RadiobuttonCheckedSvg":{"Option":"radiobutton-checked-svg","Value":""},"RadiobuttonSvg":{"Option":"radiobutton-svg","Value":""},"RunScript":{"Option":"run-script","Value":null},"SslCrtPath":{"Option":"ssl-crt-path","Value":""},"SslKeyPassword":{"Option":"ssl-key-password","Value":""},"SslKeyPath":{"Option":"ssl-key-path","Value":""},"Username":{"Option":"username","Value":""},"UserStyleSheet":{"Option":"user-style-sheet","Value":""},"ViewportSize":{"Option":"viewport-size","Value":""},"WindowStatus":{"Option":"window-status","Value":""},"Zoom":{"Option":"zoom","IsSet":false,"Value":0},"FooterCenter":{"Option":"footer-center","Value":""},"FooterFontName":{"Option":"footer-font-name","Value":""},"FooterFontSize":{"Option":"footer-font-size","IsSet":false,"Value":0},"FooterHTML":{"Option":"footer-html","Value":""},"FooterLeft":{"Option":"footer-left","Value":""},"FooterLine":{"Option":"footer-line","Value":false},"FooterRight":{"Option":"footer-right","Value":""},"FooterSpacing":{"Option":"footer-spacing","IsSet":false,"Value":0},"HeaderCenter":{"Option":"header-center","Value":""},"HeaderFontName":{"Option":"header-font-name","Value":""},"HeaderFontSize":{"Option":"header-font-size","IsSet":false,"Value":0},"HeaderHTML":{"Option":"header-html","Value":""},"HeaderLeft":{"Option":"header-left","Value":""},"HeaderLine":{"Option":"header-line","Value":false},"HeaderRight":{"Option":"header-right","Value":""},"HeaderSpacing":{"Option":"header-spacing","IsSet":false,"Value":0},"Replace":{"Option":"replace","Value":null}},"InputFile":"-","Base64PageData":"PCFkb2N0eXBlIGh0bWw+PGh0bWw+PGhlYWQ+PHRpdGxlPldLSFRNTFRPUERGIFRFU1Q8L3RpdGxlPjwvaGVhZD48Ym9keT5IRUxMTyBQREY8L2JvZHk+PC9odG1sPg=="}]}`
	if want != string(jb) {
		t.Errorf("Want JSON:\n%s\nHave:\n%s", want, string(jb))
	}
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

	want := wantArgString()
	if pdfgFromJSON.ArgString() != want {
		t.Errorf("Want argstring:\n%s\nHave:\n%s", want, pdfgFromJSON.ArgString())
	}

	pdfgFromJSON.OutputFile = "./testfiles/TestNewPDFGeneratorFromJSON.PDF"

	err = pdfgFromJSON.Create()
	if err != nil {
		t.Fatal(err)
	}

}

func TestNewPDFGeneratorFromJSONWithReader(t *testing.T) {

	pdfg := NewPDFPreparer()
	htmlfile, err := ioutil.ReadFile("./testfiles/htmlsimple.html")
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
