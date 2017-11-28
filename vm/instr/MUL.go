package instr

import (
	"fmt"

	"../ivm"
)

// MUL multiplies the content of two source registers into the destination register
type MUL struct {
	args ivm.InstructionArgsArithmetic
}

// Execute runs the given MUL instruction
func (i MUL) Execute(ip ivm.InstructionProxy) {
	source1 := ip.RegisterInt32(i.args.Source1)
	source2 := ip.RegisterInt32(i.args.Source2)
	ip.SetRegisterInt32(i.args.Destination, source1*source2)
}

// Assembly returns the representation in assembly language
func (i MUL) Assembly() string {
	return fmt.Sprintf("MUL %s", i.args.ASM())
}

// MakeMUL makes a MUL instruction for the given args
func MakeMUL(args ivm.InstructionArgs) ivm.Instruction {
	return MUL{args: args.ArithmeticFormat()}
}
