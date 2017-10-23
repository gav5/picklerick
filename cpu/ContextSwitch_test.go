package cpu

import (
	"reflect"
	"testing"

	"../proc"
	"../prog"
	"../reg"
)

func TestContextSwitch(t *testing.T) {
	c := State{}
	pcb1 := proc.PCB{
		ProgramCounter: 0x2,
		Registers: reg.List{
			/* 00 */ 0xFEEDFACE, 0xDEADBEEF, 0xFFFFFFFF, 0xFFFFFFFF,
			/* 04 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
			/* 08 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
			/* 12 */ 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF,
		},
		Program: prog.Program{
			Job: prog.Job{
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
		},
	}
	c.ContextSwitch(pcb1)
	if !reflect.DeepEqual(c.Registers, pcb1.Registers) {
		t.Errorf("registers list should be equal:\n  -> PCB %v\n  -> CPU %v\n", pcb1.Registers, c.Registers)
	}
	if !reflect.DeepEqual(c.ProgramCounter, pcb1.ProgramCounter) {
		t.Errorf("program counters should be equal:\n  -> PCB %v\n  -> CPU %v\n", pcb1.ProgramCounter, c.ProgramCounter)
	}
	if !reflect.DeepEqual(c.Program, pcb1.Program) {
		t.Errorf("programs should be the same:\n  -> PCB %v\n  -> CPU %v\n", pcb1.Program, c.Program)
	}
}
