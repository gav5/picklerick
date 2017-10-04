package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// MULI multiplies a data value directly with the content of a register
type MULI struct {
	args instrType.ArgsBranch
}

// Exec runs the given MULI instruction
func (i MULI) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i MULI) ASM() string {
	return fmt.Sprintf("MULI %s", i.args.ASM())
}

// MakeMULI makes a MULI instruction for the given args
func MakeMULI(args instrType.Args) instrType.Base {
	return MULI{args: args.BranchFormat()}
}
