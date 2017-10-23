package instr

import (
	"../cpuType"
	"../instrType"
)

// NOP does nothing and moves to the next instruction
type NOP struct {
	// no arguments needed
}

// Exec runs the given NOP instruction
func (i NOP) Exec(state *cpuType.State) {
	// this does nothing, so do nothing!
}

// ASM returns the representation in assembly language
func (i NOP) ASM() string {
	return "NOP"
}

// MakeNOP makes a NOP instruction for the given args
func MakeNOP(args instrType.Args) instrType.Base {
	return NOP{}
}
