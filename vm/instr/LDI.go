package instr

import (
	"fmt"

	"../ivm"
)

// LDI loads some data/address directly to the contents of a register
type LDI struct {
	args ivm.InstructionArgsBranch
}

// Execute runs the given LDI instruction
func (i LDI) Execute(ip ivm.InstructionProxy) {
	addr := i.args.Address
	if addr == 0 {
		base := ip.RegisterWord(i.args.Base)
		ip.SetRegisterWord(i.args.Destination, base)
	} else {
		ip.SetRegisterWord(i.args.Destination, ivm.Word(addr))
	}
}

// Assembly returns the representation in assembly language
func (i LDI) Assembly() string {
	return fmt.Sprintf("LDI %s", i.args.ASMHex())
}

// MakeLDI makes an LDI instruction for the given args
func MakeLDI(args ivm.InstructionArgs) ivm.Instruction {
	return LDI{args: args.BranchFormat()}
}
