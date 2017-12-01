package instr

import (
	"fmt"

	"../ivm"
)

// ADDI adds a data value directly to the contents of a register
type ADDI struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the ADDI instruction
func (i ADDI) Execute(ip ivm.InstructionProxy) {
	ival := ivm.Word(i.args.Address).Int32()
	dest := ip.RegisterInt32(i.args.Destination)
	ip.SetRegisterInt32(i.args.Destination, dest+ival)
}

// Assembly returns the representation in assembly language
func (i ADDI) Assembly() string {
	return fmt.Sprintf("ADDI %s", i.args.ASM())
}

// MakeADDI makes an ADDI instruction for the given args
func MakeADDI(args ivm.InstructionArgs) ivm.Instruction {
	return ADDI{args: args.BranchFormat()}
}
