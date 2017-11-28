package instr

import (
	"../ivm"
)

// HLT signals the logical end of a program (i.e. it "halts" the program)
type HLT struct {
}

// Execute runs the HLT instruction
func (i HLT) Execute(ip ivm.InstructionProxy) {
	ip.Halt()
}

// Assembly returns the representation in assembly language
func (i HLT) Assembly() string {
	return "HLT"
}

// MakeHLT makes an HLT instruction for the given args
func MakeHLT(args ivm.InstructionArgs) ivm.Instruction {
	return HLT{}
}
