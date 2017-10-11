package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// ST stores the content of a register into an address
type ST struct {
	args instrType.ArgsBranch
}

// Exec runs the given ST instruction
func (i ST) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i ST) ASM() string {
	if i.args.Base == 0x0 {
		return fmt.Sprintf("ST %s %s", i.args.Destination.ASM(), i.args.Address.Dec())
	}
	return fmt.Sprintf("ST (%s) %s", i.args.Destination.ASM(), i.args.Base.ASM())
}

// MakeST makes an ST instruction for the given args
func MakeST(args instrType.Args) instrType.Base {
	return ST{args: args.BranchFormat()}
}
