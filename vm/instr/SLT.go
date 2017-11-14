package instr

import (
	"fmt"

	"../ivm"
	
)

// SLT sets the destination register to 1 if the first source register is less than the
// branch register; otherwise, it sets the destination register to 0
type SLT struct {
	args ivm.InstructionArgsArithmetic
}

// Execute runs the given SLT instruction
func (i SLT) Execute(ip ivm.InstructionProxy) {
	// TODO: make this actually do what it's supposed to do
}

// Assembly returns the representation in assembly language
func (i SLT) Assembly() string {
	return fmt.Sprintf("SLT %s %s %s", i.args.Destination.ASM(), i.args.Source1.ASM(), i.args.Source2.ASM())
}

// MakeSLT makes an SLT instruction for the given args
func MakeSLT(args ivm.InstructionArgs) ivm.Instruction {
	return SLT{args: args.ArithmeticFormat()}
}