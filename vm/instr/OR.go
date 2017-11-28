package instr

import (
	"fmt"

	"../ivm"
)

// OR logical OR's the contents of the two source registers into a destination register
type OR struct {
	args ivm.InstructionArgsArithmetic
}

// Execute runs the given OR instruction
func (i OR) Execute(ip ivm.InstructionProxy) {
	source1 := ip.RegisterBool(i.args.Source1)
	source2 := ip.RegisterBool(i.args.Source2)
	ip.SetRegisterBool(i.args.Destination, source1 || source2)
}

// Assembly returns the representation in assembly language
func (i OR) Assembly() string {
	return fmt.Sprintf("OR %s", i.args.ASM())
}

// MakeOR makes an OR instruction for the given args
func MakeOR(args ivm.InstructionArgs) ivm.Instruction {
	return OR{args: args.ArithmeticFormat()}
}
