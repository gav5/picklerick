package instr

import (
	"fmt"

	"../ivm"
)

// BNE branches to an address when the content of the branch and destination registers
// are not equal to one another
type BNE struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the BNE instruction
func (i BNE) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterInt32(i.args.Base)
	dest := ip.RegisterInt32(i.args.Destination)
	if base != dest {
		ip.SetProgramCounter(i.args.Address)
	}
}

// Assembly returns the representation in assembly language
func (i BNE) Assembly() string {
	if i.args.Base == 0x0 {
		return fmt.Sprintf("BNE %s %s", i.args.Destination.ASM(), i.args.Address.Dec())
	}
	return fmt.Sprintf("BNE %s %s 0x%04X", i.args.Base.ASM(), i.args.Destination.ASM(), uint32(i.args.Address))
}

// MakeBNE makes a BNE instruction for the given args
func MakeBNE(args ivm.InstructionArgs) ivm.Instruction {
	return BNE{args: args.BranchFormat()}
}
