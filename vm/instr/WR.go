package instr

import (
	"fmt"

	"../ivm"
)

// WR writes the content of the accumulator into the O/P buffer
type WR struct {
	args ivm.InstructionArgsIO
}

// Execute runs the given WR instruction
func (i WR) Execute(ip ivm.InstructionProxy) {
	r1 := ip.RegisterWord(i.args.Register1)
	ip.AddressWriteWord(i.args.Address, r1)
}

// Assembly returns the representation in assembly language
func (i WR) Assembly() string {
	return fmt.Sprintf("WR %s", i.args.ASM())
}

// MakeWR makes a WR instruction for the given args
func MakeWR(args ivm.InstructionArgs) ivm.Instruction {
	return WR{args: args.IOFormat()}
}
