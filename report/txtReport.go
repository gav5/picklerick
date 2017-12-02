package report

import (
	"fmt"
	"io"

	"../vm/ivm"
)

// txtReport describes a report in plain text
// (based on the base reportBase interface, which is more generic)
type txtReport interface {
	reportBase
	Title() string
	Fprint(w io.Writer) error
}

type txtReportBuilder struct{}

// txt reports are saves as type *.txt for use as a text file
func (txtReportBuilder) Extension() string {
	return "txt"
}

// txt reports are built using plain text in variable format
// for this reason, it's best to defer to the report for formatting
// (a custom writer type is loaded in for convenience)
func (txtReportBuilder) Fprint(w io.Writer, rRAW reportBase) error {
	var err error

	r := rRAW.(txtReport)

	// print the title of the report
	_, err = fmt.Fprintf(w, "%s\n", r.Title())
	if err != nil {
		return err
	}

	// defer to the custom printing function
	err = r.Fprint(w)
	if err != nil {
		return err
	}

	// print a final newline
	_, err = fmt.Fprintln(w)
	if err != nil {
		return err
	}

	return nil
}

func fprintHeader(w io.Writer, headerTitle string) error {
	_, err := fmt.Fprintf(w, "\n\n%s:", headerTitle)
	return err
}

func fprintProperty(w io.Writer, pName string, pVal interface{}) error {
	_, err := fmt.Fprintf(w, "\n%s: %v", pName, pVal)
	return err
}

func fprintWords(w io.Writer, wordsArray []ivm.Word) error {
	for i, word := range wordsArray {
		var sep string
		if i%4 > 0 {
			sep = " "
		} else {
			sep = "\n"
		}
		_, err := fmt.Fprintf(w, "%s0x%08X", sep, uint32(word))
		if err != nil {
			return err
		}
	}
	return nil
}

func fprintBuffer(w io.Writer, wordsArray []ivm.Word, offset int) error {
	for i, word := range wordsArray {
		addr := ivm.AddressForIndex(offset + i)
		var sep string
		if i%4 > 0 {
			sep = "  "
		} else {
			fn, _ := addr.FramePair()
			sep = fmt.Sprintf("\n%02X  ", uint32(fn))
		}
		_, err := fmt.Fprintf(
			w, "%s[%03X: 0x%08X]",
			sep, uint32(addr), uint32(word),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
