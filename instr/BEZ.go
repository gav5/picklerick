package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// BEZ branches to an address when the contents of the branch register equals zero
type BEZ struct {
	args instrType.ArgsBranch
}

// Exec run the BEZ instruction
func (i BEZ) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i BEZ) ASM() string {
	return fmt.Sprintf("BEZ %s", i.args.ASM())
}

// MakeBEZ makes the BEZ instruction for the given args
func MakeBEZ(args instrType.Args) instrType.Base {
	return BEZ{args: args.BranchFormat()}
}
