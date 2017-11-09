package instr

import (
	"fmt"

	"../ivm"
)

// ADD adds the contents of the two source registers into a destination register
type ADD struct {
	args ivm.InstructionArgsArithmetic
}

// Execute runs the ADD instruction
func (i ADD) Execute(ip ivm.InstructionProxy) {
	source1 := ip.RegisterInt32(i.args.Source1)
	source2 := ip.RegisterInt32(i.args.Source2)
	ip.SetRegisterInt32(i.args.Destination, source1+source2)
}

// Assembly returns the representation in assembly language
func (i ADD) Assembly() string {
	return fmt.Sprintf("ADD %s", i.args.ASM())
}

// MakeADD makes an ADD instruction for the given args
func MakeADD(args ivm.InstructionArgs) ivm.Instruction {
	return ADD{args: args.ArithmeticFormat()}
}
