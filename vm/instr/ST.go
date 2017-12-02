package instr

import (
	"fmt"

	"../ivm"
)

// ST stores the content of a register into an address
type ST struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the given ST instruction
func (i ST) Execute(ip ivm.InstructionProxy) {
	var addr ivm.Address
	if i.args.Destination == ivm.R0 {
		addr = i.args.Address
	} else {
		addr = ivm.Address(ip.RegisterWord(i.args.Destination))
	}
	contents := ip.RegisterWord(i.args.Base)
	ip.AddressWriteWord(addr, contents)
}

// Assembly returns the representation in assembly language
func (i ST) Assembly() string {
	if i.args.Destination == 0x0 {
		return fmt.Sprintf("ST %s %s", i.args.Base.ASM(), i.args.Address.Dec())
	}
	return fmt.Sprintf("ST (%s) %s", i.args.Destination.ASM(), i.args.Base.ASM())
}

// MakeST makes an ST instruction for the given args
func MakeST(args ivm.InstructionArgs) ivm.Instruction {
	return ST{args: args.BranchFormat()}
}
