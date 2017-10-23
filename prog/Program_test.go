package prog

import (
	"reflect"
	"testing"
)

func TestProgramGetWords(t *testing.T) {
	p1 := Program{
		Job: Job{
			ID:             0x01,
			NumberOfWords:  0x04,
			PriorityNumber: 0x01,
			Instructions: []uint32{
				0xFEEDFACE,
				0xF0F0F0F0,
				0xDEADBEEF,
				0x00000000,
			},
		},
		Data: Data{
			InputBufferSize:  0x05,
			OutputBufferSize: 0x02,
			TempBufferSize:   0x03,
			DataBlock: [44]uint32{
				/* 00 */ 0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 04 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 08 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 12 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 16 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 20 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 24 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 28 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 32 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 36 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
				/* 40 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xF0F0F0F0,
			},
		},
	}
	words, err := p1.GetWords()
	if err != nil {
		t.Errorf("error getting words for program: %v", err)
	}
	p2 := Program{}
	err = p2.SetWords(words)
	if err != nil {
		t.Errorf("error settings words for program: %v", err)
	}
	if !reflect.DeepEqual(p1, p2) {
		t.Errorf("expected p1 and p2 to be the same!\n(p1 %v)\n(p2 %v)", p1, p2)
	}
}
