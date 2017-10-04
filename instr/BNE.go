package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// BNE branches to an address when the content of the branch and destination registers
// are not equal to one another
type BNE struct {
	args instrType.ArgsBranch
}

// Exec runs the BNE instruction
func (i BNE) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i BNE) ASM() string {
	return fmt.Sprintf("BNE %s", i.args.ASM())
}

// MakeBNE makes a BNE instruction for the given args
func MakeBNE(args instrType.Args) instrType.Base {
	return BNE{args: args.BranchFormat()}
}
