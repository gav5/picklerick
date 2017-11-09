package instr

import (
	"fmt"

	"../ivm"
)

// MULI multiplies a data value directly with the content of a register
type MULI struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the given MULI instruction
func (i MULI) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterInt32(i.args.Base)
	dest := ip.RegisterInt32(i.args.Destination)
	ip.SetRegisterInt32(i.args.Destination, base*dest)
}

// Assembly returns the representation in assembly language
func (i MULI) Assembly() string {
	return fmt.Sprintf("MULI %s", i.args.ASM())
}

// MakeMULI makes a MULI instruction for the given args
func MakeMULI(args ivm.InstructionArgs) ivm.Instruction {
	return MULI{args: args.BranchFormat()}
}
