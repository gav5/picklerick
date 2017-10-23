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
	words, err := d.getWords()
	if err != nil {
		t.Errorf("error getting words for data: %v", err)
	}
	d2 := new(Data)
	err = d2.setWords(words)
	if err != nil {
		t.Errorf("error setting words for data: %v", err)
	}
	if !reflect.DeepEqual(d, *d2) {
		t.Errorf("expected d and d2 to be the same! (expected %v, got %v)", d, *d2)
	}
}
