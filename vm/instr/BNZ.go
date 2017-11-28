package instr

import (
	"fmt"

	"../ivm"
)

// BNZ branches to an address when the contents of the branch register is not zero
type BNZ struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the BNZ instruction
func (i BNZ) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterInt32(i.args.Base)
	if base != 0 {
		ip.SetProgramCounter(i.args.Address)
	}
}

// Assembly returns the representation in assembly language
func (i BNZ) Assembly() string {
	return fmt.Sprintf("BNZ %s", i.args.ASM())
}

// MakeBNZ makes a BNZ instruction for the given args
func MakeBNZ(args ivm.InstructionArgs) ivm.Instruction {
	return BNZ{args: args.BranchFormat()}
}
