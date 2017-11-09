package instr

import (
	"fmt"

	"../ivm"
)

// BLZ branches to an address when the content of the branch register is less than 0
type BLZ struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the BLZ instruction
func (i BLZ) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterInt32(i.args.Base)
	// NOTE: this should never happen because this is an unsigned value!!
	if base < 0 {
		ip.SetProgramCounter(i.args.Address)
	}
}

// Assembly returns the representation in assembly language
func (i BLZ) Assembly() string {
	return fmt.Sprintf("BLZ %s", i.args.ASM())
}

// MakeBLZ makes the BLZ instruction for the given args
func MakeBLZ(args ivm.InstructionArgs) ivm.Instruction {
	return BLZ{args: args.BranchFormat()}
}
