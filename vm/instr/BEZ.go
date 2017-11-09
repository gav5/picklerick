package instr

import (
	"fmt"

	"../ivm"
)

// BEZ branches to an address when the contents of the branch register equals zero
type BEZ struct {
	args ivm.InstructionArgsBranch
}

// Execute run the BEZ instruction
func (i BEZ) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterInt32(i.args.Base)
	if base == 0 {
		ip.SetProgramCounter(i.args.Address)
	}
}

// Assembly returns the representation in assembly language
func (i BEZ) Assembly() string {
	return fmt.Sprintf("BEZ %s", i.args.ASM())
}

// MakeBEZ makes the BEZ instruction for the given args
func MakeBEZ(args ivm.InstructionArgs) ivm.Instruction {
	return BEZ{args: args.BranchFormat()}
}
