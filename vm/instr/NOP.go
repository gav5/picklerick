package instr

import (
	"../ivm"
	
)

// NOP does nothing and moves to the next instruction
type NOP struct {
	// no arguments needed
}

// Execute runs the given NOP instruction
func (i NOP) Execute(ip ivm.InstructionProxy) {
	// this does nothing, so do nothing!
}

// Assembly returns the representation in assembly language
func (i NOP) Assembly() string {
	return "NOP"
}

// MakeNOP makes a NOP instruction for the given args
func MakeNOP(args ivm.InstructionArgs) ivm.Instruction {
	return NOP{}
}
