package instr

import (
	"fmt"

	"../instrType"
	"../proc"
)

// BLZ branches to an address when the content of the branch register is less than 0
type BLZ struct {
	args instrType.ArgsBranch
}

// Exec runs the BLZ instruction
func (i BLZ) Exec(pcb proc.PCB) proc.PCB {
	// TODO: make this actually do what it's supposed to do
	return pcb
}

// ASM returns the representation in assembly language
func (i BLZ) ASM() string {
	return fmt.Sprintf("BLZ %s", i.args.ASM())
}

// MakeBLZ makes the BLZ instruction for the given args
func MakeBLZ(args instrType.Args) instrType.Base {
	return BLZ{args: args.BranchFormat()}
}
