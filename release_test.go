package wkhtmltopdf

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

// flags which are defaults are not present in the code
var ignoredFlags = []string{
	"top",
	"enable-javascript",
	"background",
	"no-debug-javascript",
	"enable-external-links",
	"outline",
	"disable-toc-back-links",
	"no-footer-line",
	"no-header-line",
	"no-print-media-type",
	"include-in-outline",
	"enable-smart-shrinking",
	"resolve-relative-links",
	"collate",
	"disable-forms",
	"disable-plugins",
	"stop-slow-scripts",
	"images",
	"enable-internal-links",
}

// TestIfArgsAreIncluded checks if all args in the help output of wkhtmltopdf are present in the code.
// If checks if the arguments are in any option set, it currently is not smart enough to check if they are
// in the correct options set (like TOC options, page options etc.).
// This test is skipped by default.
func TestIfArgsAreIncluded(t *testing.T) {
	t.SkipNow() // remove if you want to run this test

	// run wkhtmltopdf with extended-help argument
	wkpdf, err := NewPDFGenerator()
	if err != nil {
		t.Fatal(err)
	}
	wkpdf.ExtendedHelp.Set(true)

	err = wkpdf.Create()
	if err != nil {
		t.Fatal(err)
	}

	// cut off the buffer after "Page sizes:', this is where the arguments end
	bufStr := wkpdf.Buffer().String()
	bufStr = bufStr[:strings.Index(bufStr, "Page sizes:")]

	// use simple regex to get all long arument flags
	argRegex, err := regexp.Compile(`--[a-z-]+\s`)
	if err != nil {
		t.Fatal(err)
	}
	helpFlags := argRegex.FindAllString(bufStr, -1)

	// put the flags in a map to remove duplicates
	mHelpFlags := make(map[string]bool)
	for _, flag := range helpFlags {
		mHelpFlags[strings.TrimSpace(strings.TrimPrefix(flag, "--"))] = true
	}
	t.Logf("found %d arguments in extended help", len(mHelpFlags))

	var unusedFlags []string

	// for each option, check if it is in the code
	global := newGlobalOptions()
	headerFooter := newHeaderAndFooterOptions()
	outline := newOutlineOptions()
	page := newPageOptions()
	toc := newTocOptions()
	for flag, _ := range mHelpFlags {
		// check if flag is ignored
		ignored := false
		for _, ignoredFlag := range ignoredFlags {
			if flag == ignoredFlag {
				ignored = true
				continue
			}
		}
		if ignored {
			continue
		}

		// check if flag is present
		found := false
		found = optionsHaveArg(global, flag)
		if !found {
			found = optionsHaveArg(headerFooter, flag)
		}
		if !found {
			found = optionsHaveArg(outline, flag)
		}
		if !found {
			found = optionsHaveArg(page, flag)
		}
		if !found {
			found = optionsHaveArg(toc, flag)
		}
		if !found {
			unusedFlags = append(unusedFlags, flag)
		}
	}

	if len(unusedFlags) > 0 {
		t.Errorf("%d unused flags:\n%s", len(unusedFlags), strings.Join(unusedFlags, "\n"))
	}
}

func optionsHaveArg(opts interface{}, arg string) bool {
	rv := reflect.Indirect(reflect.ValueOf(opts))
	if rv.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < rv.NumField(); i++ {
		optionField := rv.Field(i).FieldByName("option")
		if optionField.IsZero() {
			continue
		}
		if optionField.String() == arg {
			return true
		}
	}
	return false
}
