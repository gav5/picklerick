package util

import (
	"testing"
)

var extract8Tests = []struct {
	in  string
	val uint8
	err error
}{
	{"", 0, HexExtractionFormatError{in: ""}},
	{"0", 0, nil},
	{"01", 1, nil},
	{"ff", 0xff, nil},
	{"fe", 0xfe, nil},
}

func TestHexExtract8(t *testing.T) {
	for _, tt := range extract8Tests {
		val, err := HexExtract8(tt.in)
		if (err != tt.err) || (val != tt.val) {
			t.Errorf("HexExtract8(%q) => %#02X,%v; want %#02X,%v", tt.in, val, err, tt.val, tt.err)
		}
	}
}
