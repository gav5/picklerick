package prog

import (
	"reflect"
	"testing"
)

func TestJobWordTranslation(t *testing.T) {
	j := Job{
		ID:             0x01,
		NumberOfWords:  0x04,
		PriorityNumber: 0x01,
		Instructions: []uint32{
			0xFEEDFACE,
			0x00000000,
			0xDEADBEEF,
			0x00000000,
		},
	}
	words := j.getWords()
	words2 := []uint32{0xFEEDFACE, 0x00000000, 0xDEADBEEF, 0x00000000}
	if !reflect.DeepEqual(words, words2) {
		t.Errorf("expected words and words2 to be the same! (expected %v, got %v)", words, words2)
	}
}
