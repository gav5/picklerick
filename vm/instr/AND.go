package instr

import (
	"fmt"

	"../ivm"
)

// AND logical AND's the contents of two source registers into a desgination register
type AND struct {
	args ivm.InstructionArgsArithmetic
}

// Execute runs the AND instruction
func (i AND) Execute(ip ivm.InstructionProxy) {
	source1 := ip.RegisterBool(i.args.Source1)
	source2 := ip.RegisterBool(i.args.Source2)
	ip.SetRegisterBool(i.args.Destination, source1 && source2)
}

// Assembly returns the representation in assembly language
func (i AND) Assembly() string {
	return fmt.Sprintf("AND %s", i.args.ASM())
}

// MakeAND makes an AND instruction for the given args
func MakeAND(args ivm.InstructionArgs) ivm.Instruction {
	return AND{args: args.ArithmeticFormat()}
}
