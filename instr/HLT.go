package instr

import (
	"../instrType"
	"../proc"
)

// HLT signals the logical end of a program (i.e. it "halts" the program)
type HLT struct {
}

// Exec runs the HLT instruction
func (i HLT) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i HLT) ASM() string {
	return "HLT"
}

// MakeHLT makes an HLT instruction for the given args
func MakeHLT(args instrType.Args) instrType.Base {
	return HLT{}
}
