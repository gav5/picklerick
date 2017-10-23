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
	words, err := j.getWords()
	if err != nil {
		t.Errorf("error getting words for job: %v", err)
	}
	j2 := Job{}
	err = (&j2).setWords(words)
	if err != nil {
		t.Errorf("error setting words for job: %v", err)
	}
	if !reflect.DeepEqual(j, j2) {
		t.Errorf("expected j and j2 to be the same! (expected %v, got %v)", j, j2)
	}
}
