package instr

import (
	"../cpuType"
	"../instrType"
)

// HLT signals the logical end of a program (i.e. it "halts" the program)
type HLT struct {
}

// Exec runs the HLT instruction
func (i HLT) Exec(state *cpuType.State) {
	state.Halt()
}

// ASM returns the representation in assembly language
func (i HLT) ASM() string {
	return "HLT"
}

// MakeHLT makes an HLT instruction for the given args
func MakeHLT(args instrType.Args) instrType.Base {
	return HLT{}
}
