package disk

import (
	"reflect"
	"testing"

	"../prog"
)

func TestProgramStorage(t *testing.T) {
	p1 := prog.Program{
		Job: prog.Job{
			ID:             0x01,
			NumberOfWords:  0x04,
			PriorityNumber: 0x01,
			Instructions: []uint32{
				0xFEEDFACE,
				0xF0F0F0F0,
				0xDEADBEEF,
				0x11111111,
			},
		},
		Data: prog.Data{
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
	if err := StoreProgram(p1); err != nil {
		t.Errorf("cannot store program because of error: %v", err)
	}
	p2 := prog.Program{}
	if err := LoadProgram(&p2, p1.Job.ID); err != nil {
		t.Errorf("cannot load program because of error: %v", err)
	}
	if !reflect.DeepEqual(p1, p2) {
		t.Errorf("expected p1 and p2 to be the same!\n(p1 %v)\n(p2 %v)", p1, p2)
	}
}
