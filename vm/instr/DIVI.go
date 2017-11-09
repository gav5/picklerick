package instr

import (
	"fmt"

	"../ivm"
)

// DIVI divides a data value directly with the content of a register
type DIVI struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the DIVI instruction
func (i DIVI) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterInt32(i.args.Base)
	dest := ip.RegisterInt32(i.args.Destination)
	ip.SetRegisterInt32(i.args.Destination, base/dest)
}

// Assembly returns the representation in assembly language
func (i DIVI) Assembly() string {
	return fmt.Sprintf("DIVI %s", i.args.ASM())
}

// MakeDIVI makes a DIVI instruction for the given args
func MakeDIVI(args ivm.InstructionArgs) ivm.Instruction {
	return DIVI{args: args.BranchFormat()}
}
