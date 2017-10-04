package instr

import (
	"../instrType"
	"../proc"
)

// NOP does nothing and moves to the next instruction
type NOP struct {
	// no arguments needed
}

// Exec runs the given NOP instruction
func (i NOP) Exec(pcb proc.PCB) proc.PCB {
	// this does nothing, so return the same state
	return pcb
}

// ASM returns the representation in assembly language
func (i NOP) ASM() string {
	return "NOP"
}

// MakeNOP makes a NOP instruction for the given args
func MakeNOP(args instrType.Args) instrType.Base {
	return NOP{}
}
