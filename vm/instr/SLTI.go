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
	source1 := ip.RegisterBool(i.args.Source1)
	source2 := ip.RegisterBool(i.args.Source2)
	if(source1<source2){
		ip.SetRegisterBool(i.args.Destination, 1)
	} else {
		ip.SetRegisterBool(i.args.Destination, 0)
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
