package instr

import (
	"fmt"

	"../ivm"
)

// SUB subtracts the contents of the two source registers into the destination register
type SUB struct {
	args ivm.InstructionArgsArithmetic
}

// Execute runs the given SUB instruction
func (i SUB) Execute(ip ivm.InstructionProxy) {
	source1 := ip.RegisterInt32(i.args.Source1)
	source2 := ip.RegisterInt32(i.args.Source2)
	ip.SetRegisterInt32(i.args.Destination, source1-source2)
}

// Assembly returns the representation in assembly language
func (i SUB) Assembly() string {
	return fmt.Sprintf("SUB %s", i.args.ASM())
}

// MakeSUB makes a SUB instruction for the given args
func MakeSUB(args ivm.InstructionArgs) ivm.Instruction {
	return SUB{args: args.ArithmeticFormat()}
}
