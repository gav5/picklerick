package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// BEQ branches to an address when the contents of the branch and destination registers
// are equal to one another
type BEQ struct {
	args instrType.ArgsBranch
}

// Exec runs the BEQ instruction
func (i BEQ) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i BEQ) ASM() string {
	return fmt.Sprintf("BEQ %s", i.args.ASM())
}

// MakeBEQ makes the BEQ instruction for the given args
func MakeBEQ(args instrType.Args) instrType.Base {
	return BEQ{args: args.BranchFormat()}
}
