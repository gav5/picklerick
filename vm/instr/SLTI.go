package instr

import (
	"fmt"

	"../ivm"

)

// SLTI sets the destination register to 1 if the first source register is less than a
// data value; otherwise, it sets the destination register to 0
type SLTI struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the given SLTI instruction
func (i SLTI) Execute(ip ivm.InstructionProxy) {
	base := ip.RegisterWord(i.args.Base)
	data := ivm.Word(i.args.Address)
	if base < data {
		ip.SetRegisterBool(i.args.Destination, true)
	} else {
		ip.SetRegisterBool(i.args.Destination, false)
	}
}

// Assembly returns the representation in assembly language
func (i SLTI) Assembly() string {
	return fmt.Sprintf("SLTI %s", i.args.ASM())
}

// MakeSLTI makes an SLTI instruction for the given args
func MakeSLTI(args ivm.InstructionArgs) ivm.Instruction {
	return SLTI{args: args.BranchFormat()}
}
