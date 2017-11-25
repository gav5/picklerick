package instr

import (
	"fmt"

	"../ivm"

)

// MOVI transfers address/data directly into a register
type MOVI struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the given MOVI instruction
func (i MOVI) Execute(ip ivm.InstructionProxy) {
	address := i.args.Address	// get source address
	source := ip.AddressFetchWord(address) // fetch source
	ip.SetRegisterWord(i.args.Destination, source)
}

// Assembly returns the representation in assembly language
func (i MOVI) Assembly() string {
	return fmt.Sprintf("MOVI %s", i.args.ASM())
}

// MakeMOVI makes a MOVI instruction for the given args
func MakeMOVI(args ivm.InstructionArgs) ivm.Instruction {
	return MOVI{args: args.BranchFormat()}
}
