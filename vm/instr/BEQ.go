package instr

import (
	"fmt"

	"../ivm"
)

// BEQ branches to an address when the contents of the branch and destination registers
// are equal to one another
type BEQ struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the BEQ instruction
func (i BEQ) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterInt32(i.args.Base)
	dest := ip.RegisterInt32(i.args.Destination)
	if base == dest {
		ip.SetProgramCounter(i.args.Address)
	}
}

// Assembly returns the representation in assembly language
func (i BEQ) Assembly() string {
	return fmt.Sprintf("BEQ %s", i.args.ASM())
}

// MakeBEQ makes the BEQ instruction for the given args
func MakeBEQ(args ivm.InstructionArgs) ivm.Instruction {
	return BEQ{args: args.BranchFormat()}
}
