package instr

import (
	"fmt"

	"../ivm"

)

// LW loads the content of an address into a register
type LW struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the given LW instruction
func (i LW) Execute(ip ivm.InstructionProxy) {
	contents := ip.AddressFetchWord(i.args.Address)
	ip.SetRegisterWord(i.args.Destination, contents)
}

// Assembly returns the representation in assembly language
func (i LW) Assembly() string {
	return fmt.Sprintf("LW %s %s(%s)", i.args.Destination.ASM(), i.args.Address.Dec(), i.args.Base.ASM())
}

// MakeLW makes an LW instruction for the given args
func MakeLW(args ivm.InstructionArgs) ivm.Instruction {
	return LW{args: args.BranchFormat()}
}
