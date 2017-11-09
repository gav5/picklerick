package instr

import (
	"fmt"

	"../ivm"
)

// BGZ branches to an address when the contents of the branch register is greater than 0
type BGZ struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the BGZ instruction
func (i BGZ) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterInt32(i.args.Base)
	if base > 0 {
		ip.SetProgramCounter(i.args.Address)
	}
}

// Assembly returns the representation in assembly language
func (i BGZ) Assembly() string {
	return fmt.Sprintf("BGZ %s", i.args.ASM())
}

// MakeBGZ make the BGZ instruction for the given args
func MakeBGZ(args ivm.InstructionArgs) ivm.Instruction {
	return BGZ{args: args.BranchFormat()}
}
