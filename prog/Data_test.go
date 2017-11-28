package prog

import (
	"reflect"
	"testing"
)

func TestDataWordTranslation(t *testing.T) {
	d := Data{
		InputBufferSize:  0x05,
		OutputBufferSize: 0x02,
		TempBufferSize:   0x03,
		DataBlock:        [44]uint32{0xFEEDFACE, 0xDEADBEEF},
	}
	words := d.getWords()
	words2 := [44]uint32{0xFEEDFACE, 0xDEADBEEF}
	if !reflect.DeepEqual(words, words2) {
		t.Errorf("expected words and words2 to be the same! (expected %v, got %v)", words, words2)
	}
}
