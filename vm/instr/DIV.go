package instr

import (
	"fmt"

	"../ivm"
)

// DIV divides the content of two source registers into the destination register
type DIV struct {
	args ivm.InstructionArgsArithmetic
}

// Execute runs the DIV instruction
func (i DIV) Execute(ip ivm.InstructionProxy) {
	source1 := ip.RegisterInt32(i.args.Source1)
	source2 := ip.RegisterInt32(i.args.Source2)
	ip.SetRegisterInt32(i.args.Destination, source1/source2)
}

// Assembly returns the representation in assembly language
func (i DIV) Assembly() string {
	return fmt.Sprintf("DIV %s", i.args.ASM())
}

// MakeDIV makes a DIV instruction for the given arguments
func MakeDIV(args ivm.InstructionArgs) ivm.Instruction {
	return DIV{args: args.ArithmeticFormat()}
}
